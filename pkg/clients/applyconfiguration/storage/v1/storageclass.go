// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	applyconfigurationcorev1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// StorageClassApplyConfiguration represents an declarative configuration of the StorageClass type for use
// with apply.
type StorageClassApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Provisioner                      *string                                                           `json:"provisioner,omitempty"`
	Parameters                       map[string]string                                                 `json:"parameters,omitempty"`
	ReclaimPolicy                    *corev1.PersistentVolumeReclaimPolicy                             `json:"reclaimPolicy,omitempty"`
	MountOptions                     []string                                                          `json:"mountOptions,omitempty"`
	AllowVolumeExpansion             *bool                                                             `json:"allowVolumeExpansion,omitempty"`
	VolumeBindingMode                *storagev1.VolumeBindingMode                                      `json:"volumeBindingMode,omitempty"`
	AllowedTopologies                []applyconfigurationcorev1.TopologySelectorTermApplyConfiguration `json:"allowedTopologies,omitempty"`
}

// StorageClass constructs an declarative configuration of the StorageClass type for use with
// apply.
func StorageClass(name string) *StorageClassApplyConfiguration {
	b := &StorageClassApplyConfiguration{}
	b.WithName(name)
	b.WithKind("StorageClass")
	b.WithAPIVersion("storage.k8s.io/v1")
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithKind(value string) *StorageClassApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithAPIVersion(value string) *StorageClassApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithName(value string) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithGenerateName(value string) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithNamespace(value string) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithUID(value types.UID) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithResourceVersion(value string) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithGeneration(value int64) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithCreationTimestamp(value metav1.Time) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *StorageClassApplyConfiguration) WithLabels(entries map[string]string) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *StorageClassApplyConfiguration) WithAnnotations(entries map[string]string) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *StorageClassApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *StorageClassApplyConfiguration) WithFinalizers(values ...string) *StorageClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *StorageClassApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithProvisioner sets the Provisioner field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Provisioner field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithProvisioner(value string) *StorageClassApplyConfiguration {
	b.Provisioner = &value
	return b
}

// WithParameters puts the entries into the Parameters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Parameters field,
// overwriting an existing map entries in Parameters field with the same key.
func (b *StorageClassApplyConfiguration) WithParameters(entries map[string]string) *StorageClassApplyConfiguration {
	if b.Parameters == nil && len(entries) > 0 {
		b.Parameters = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Parameters[k] = v
	}
	return b
}

// WithReclaimPolicy sets the ReclaimPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReclaimPolicy field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithReclaimPolicy(value corev1.PersistentVolumeReclaimPolicy) *StorageClassApplyConfiguration {
	b.ReclaimPolicy = &value
	return b
}

// WithMountOptions adds the given value to the MountOptions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MountOptions field.
func (b *StorageClassApplyConfiguration) WithMountOptions(values ...string) *StorageClassApplyConfiguration {
	for i := range values {
		b.MountOptions = append(b.MountOptions, values[i])
	}
	return b
}

// WithAllowVolumeExpansion sets the AllowVolumeExpansion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowVolumeExpansion field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithAllowVolumeExpansion(value bool) *StorageClassApplyConfiguration {
	b.AllowVolumeExpansion = &value
	return b
}

// WithVolumeBindingMode sets the VolumeBindingMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeBindingMode field is set to the value of the last call.
func (b *StorageClassApplyConfiguration) WithVolumeBindingMode(value storagev1.VolumeBindingMode) *StorageClassApplyConfiguration {
	b.VolumeBindingMode = &value
	return b
}

// WithAllowedTopologies adds the given value to the AllowedTopologies field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedTopologies field.
func (b *StorageClassApplyConfiguration) WithAllowedTopologies(values ...*applyconfigurationcorev1.TopologySelectorTermApplyConfiguration) *StorageClassApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAllowedTopologies")
		}
		b.AllowedTopologies = append(b.AllowedTopologies, *values[i])
	}
	return b
}