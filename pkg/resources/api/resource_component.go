package api

import (
	"reflect"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/mohae/deepcopy"
	"github.com/seal-io/utils/json"
	"k8s.io/klog/v2"
)

type ResourceComponentTerraformChange struct {
	*tfjson.ResourceChange `json:",inline"`

	Change *TerraformChange `json:"change"`
}

type TerraformChange struct {
	*tfjson.Change `json:",inline"`

	Type string `json:"type"`
}

const (
	ResourceComponentChangeTypeCreate   = "create"
	ResourceComponentChangeTypeUpdate   = "update"
	ResourceComponentChangeTypeDelete   = "delete"
	ResourceComponentChangeTypeNoChange = "no-change"
)

// Process parses the change type from the actions and sets the type to the change.
func (c *TerraformChange) Process() *TerraformChange {
	logger := klog.Background().WithName("component-change")

	switch {
	case c.Actions.Create():
		c.Type = ResourceComponentChangeTypeCreate
	case c.Actions.Update(),
		c.Actions.Replace():
		c.Type = ResourceComponentChangeTypeUpdate
	case c.Actions.Delete():
		c.Type = ResourceComponentChangeTypeDelete
	case c.Actions.NoOp(), c.Actions.Read():
		c.Type = ResourceComponentChangeTypeNoChange
	}

	var (
		diff any
		// SensitivePatchedDiff is the diff with the sensitive values patched.
		sensitivePatchedDiff any
		err                  error
	)
	if c.Change.Before != nil && c.Change.After != nil {
		// Ignore error as the diff is not critical.
		diff, err = json.CreateMergePatch(c.Change.Before, c.Change.After)
		if err != nil {
			logger.Infof("failed to create the merge diff: %v", err)
		}

		afterSensitive := deepcopy.Copy(c.Change.AfterSensitive)

		// Patch the sensitive values changes.
		sensitivePatchedDiff = patchLeaf(diff, afterSensitive, "<sensitive value(changed)>", false)
	}

	c.Change.Before = patchLeaf(c.Change.Before, c.Change.BeforeSensitive, "<sensitive value>", false)
	c.Change.After = patchLeaf(c.Change.After, c.Change.AfterSensitive, "<sensitive value>", false)

	if sensitivePatchedDiff != nil && c.Change.After != nil {
		object, err := json.PatchObject(c.Change.After, sensitivePatchedDiff)
		if err == nil {
			c.Change.After = object
		} else {
			logger.Infof("failed to patch the sensitive values: %v", err)
		}
	}

	c.Change.After = patchLeaf(c.Change.After, c.Change.AfterUnknown, "<known after apply>", true)

	c.Change.BeforeSensitive = nil
	c.Change.AfterSensitive = nil
	c.Change.AfterUnknown = nil
	c.Change.GeneratedConfig = ""
	c.Change.ReplacePaths = nil

	return c
}

// patchLeaf patch the raw value with the masked leaf value with the mask.
func patchLeaf(value, toMaskLeaf any, mask string, merge bool) any {
	logger := klog.Background().WithName("component-change")
	if value == nil || toMaskLeaf == nil {
		return value
	}

	maskedLeaf := maskLeafValues(value, toMaskLeaf, mask, merge)
	if maskedLeaf == nil {
		return value
	}

	patched, err := json.PatchObject(value, maskedLeaf)
	if err != nil {
		logger.Infof("failed to patch the leaf value: %v", err)

		return value
	}

	if ptr := reflect.ValueOf(patched); ptr.Kind() == reflect.Ptr {
		patched = ptr.Elem().Interface()
	}

	return patched
}

// maskLeafValues masks the leaf values of the raw value with the mask.
// The toMaskLeafs record the leaf key to be masked.
// If merge is true, the leaf values will be merged with raw value.
func maskLeafValues(rawValue, toMaskLeafs any, mask string, merge bool) any {
	if isEmptyValueLeaf(toMaskLeafs) {
		return nil
	}

	// If the mask value is true, replace the raw value with the mask.
	if boolVal, isBool := toMaskLeafs.(bool); isBool {
		if boolVal {
			return mask
		}
		return rawValue
	}

	if ptr := reflect.ValueOf(rawValue); ptr.Kind() == reflect.Ptr {
		rawValue = ptr.Elem().Interface()
	}

	if ptr := reflect.ValueOf(toMaskLeafs); ptr.Kind() == reflect.Ptr {
		toMaskLeafs = ptr.Elem().Interface()
	}

	switch leafVal := toMaskLeafs.(type) {
	case map[string]any:
		val, ok := rawValue.(map[string]any)
		if !ok {
			return rawValue
		}

		for k := range leafVal {
			if merge && leafVal[k] == true {
				leafVal[k] = mask
			} else {
				if _, ok := val[k]; !ok {
					if merge {
						continue
					}

					delete(leafVal, k)
					continue
				}

				leafVal[k] = maskLeafValues(val[k], leafVal[k], mask, merge)

				if leafVal[k] == nil {
					delete(leafVal, k)
				}
			}
		}

		if len(leafVal) == 0 {
			return nil
		}

		return leafVal
	case []any:
		val, ok := rawValue.([]any)
		if !ok {
			return rawValue
		}

		maskLen := len(val)
		if merge && len(leafVal) > maskLen {
			maskLen = len(leafVal)
		}

		masked := make([]any, maskLen)

		for i := range masked {
			if merge && i >= len(val) && !isEmptyValueLeaf(leafVal[i]) {
				if leafVal[i] == true {
					masked[i] = mask
					continue
				} else {
					masked[i] = leafVal[i]
					continue
				}
			}

			masked[i] = val[i]
			if i >= len(leafVal) {
				continue
			}

			m := maskLeafValues(val[i], leafVal[i], mask, merge)
			if m != nil {
				masked[i] = m
			}
		}

		return masked
	case bool:
		if leafVal {
			return mask
		}
	}

	return rawValue
}

func isEmptyValueLeaf(v any) bool {
	reflectValue := reflect.ValueOf(v)
	switch reflectValue.Kind() {
	case reflect.String, reflect.Array:
		return reflectValue.Len() == 0
	case reflect.Map, reflect.Slice:
		return reflectValue.Len() == 0 || reflectValue.IsNil()
	case reflect.Bool:
		// As the value is a boolean, it is not an empty value.
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflectValue.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return reflectValue.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return reflectValue.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return reflectValue.IsNil()
	default:
	}

	return false
}
