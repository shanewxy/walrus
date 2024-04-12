package v1

import "errors"

// EnvironmentType describes the type of environment.
// +enum
type EnvironmentType string

const (
	// EnvironmentTypeDevelopment means the environment is for development.
	EnvironmentTypeDevelopment EnvironmentType = "Development"
	// EnvironmentTypeStaging means the environment is for staging.
	EnvironmentTypeStaging EnvironmentType = "Staging"
	// EnvironmentTypeProduction means the environment is for production.
	EnvironmentTypeProduction EnvironmentType = "Production"
)

func (in EnvironmentType) String() string {
	return string(in)
}

func (in EnvironmentType) Validate() error {
	switch in {
	case EnvironmentTypeDevelopment, EnvironmentTypeStaging, EnvironmentTypeProduction:
		return nil
	default:
		return errors.New("invalid environment type")
	}
}
