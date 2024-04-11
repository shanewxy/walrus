package openapi

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/seal-io/utils/json"
	"github.com/stretchr/testify/assert"

	"github.com/seal-io/walrus/pkg/templates/api/property"
)

func TestGenSchemaDefaultPatch(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError bool
	}{
		{
			name:          "Empty",
			input:         "testdata/empty",
			expectedError: true,
		},
		{
			name:          "With simple",
			input:         "testdata/simple",
			expectedError: true,
		},
		{
			name:          "With map",
			input:         "testdata/map",
			expectedError: false,
		},
		{
			name:          "With list",
			input:         "testdata/list",
			expectedError: false,
		},
		{
			name:          "With object",
			input:         "testdata/object",
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := openapi3.NewLoader()

			ts, err := l.LoadFromFile(filepath.Join(tc.input, "schema.yaml"))
			assert.NoError(t, err)

			s := ts.Components.Schemas["variables"]
			m, err := GenSchemaDefaultPatch(context.Background(), s.Value)
			assert.NoError(t, err)

			fb, err := os.ReadFile(filepath.Join(tc.input, "expected.json"))
			assert.NoError(t, err)

			if len(fb) == 0 {
				assert.Empty(t, m)
				return
			}

			// Unmarshal and compare the result to prevent the order of the map keys from affecting the result.
			var expected, actual any
			assert.NoError(t, json.Unmarshal(fb, &expected))
			assert.NoError(t, json.Unmarshal(m, &actual))
			assert.Equal(t, expected, actual)
		})
	}
}

func TestGenSchemaDefaultWithAttribute(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError bool
	}{
		{
			name:          "Empty",
			input:         "testdata/empty",
			expectedError: true,
		},
		{
			name:          "With simple",
			input:         "testdata/simple",
			expectedError: true,
		},
		{
			name:          "With map",
			input:         "testdata/map",
			expectedError: false,
		},
		{
			name:          "With list",
			input:         "testdata/list",
			expectedError: false,
		},
		{
			name:          "With object",
			input:         "testdata/object",
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := openapi3.NewLoader()

			ts, err := l.LoadFromFile(filepath.Join(tc.input, "schema.yaml"))
			assert.NoError(t, err)

			s := ts.Components.Schemas["variables"]

			attrByte, err := os.ReadFile(filepath.Join(tc.input, "with_attributes_input.json"))
			assert.NoError(t, err)

			var attrs property.Values
			err = json.Unmarshal(attrByte, &attrs)
			assert.NoError(t, err)

			attrsDefault, err := os.ReadFile(filepath.Join(tc.input, "with_attributes_input_default.json"))
			assert.NoError(t, err)

			m, err := GenSchemaDefaultWithAttribute(context.Background(), s.Value, attrs, attrsDefault)
			assert.NoError(t, err)

			fb, err := os.ReadFile(filepath.Join(tc.input, "with_attributes_expected.json"))
			assert.NoError(t, err)

			if len(fb) == 0 {
				assert.Empty(t, m)
				return
			}

			// Unmarshal and compare the result to prevent the order of the map keys from affecting the result.
			var expected, actual any
			assert.NoError(t, json.Unmarshal(fb, &expected))
			assert.NoError(t, json.Unmarshal(m, &actual))
			assert.Equal(t, expected, actual)
		})
	}
}
