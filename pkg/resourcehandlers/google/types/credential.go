package types

import (
	"errors"
)

const (
	Project     = "project"
	Region      = "region"
	Zone        = "zone"
	Credentials = "credentials"
)

type Credential struct {
	Project     string
	Region      string
	Zone        string
	Credentials string
}

func GetCredential(configData map[string][]byte) (*Credential, error) {
	cred := &Credential{}

	cred.Project = string(configData[Project])
	if cred.Project == "" {
		return nil, errors.New("project is empty")
	}

	cred.Region = string(configData[Region])
	if cred.Region == "" {
		return nil, errors.New("region is empty")
	}

	cred.Zone = string(configData[Zone])
	if cred.Zone == "" {
		return nil, errors.New("zone is empty")
	}

	cred.Credentials = string(configData[Credentials])
	if cred.Credentials == "" {
		return nil, errors.New("credentials is empty")
	}

	return cred, nil
}
