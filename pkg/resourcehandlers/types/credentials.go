package types

import (
	"context"
	"errors"
)

type CredentialKeyType string

const (
	CredentialKey CredentialKeyType = "credential"
)

const (
	AccessKey    = "access_key"
	AccessSecret = "secret_key"
	Region       = "region"
)

type Credential struct {
	AccessKey    string
	AccessSecret string
	Region       string
}

func GetCredential(configData map[string][]byte) (*Credential, error) {
	cred := &Credential{}

	cred.AccessKey = string(configData[AccessKey])
	if cred.AccessKey == "" {
		return nil, errors.New("accessKey is empty")
	}

	cred.AccessSecret = string(configData[AccessSecret])
	if cred.AccessSecret == "" {
		return nil, errors.New("accessSecret is empty")
	}

	cred.Region = string(configData[Region])
	if cred.Region == "" {
		return nil, errors.New("region is empty")
	}

	return cred, nil
}

func CredentialFromCtx(ctx context.Context) (*Credential, error) {
	cred, ok := ctx.Value(CredentialKey).(*Credential)
	if !ok {
		return nil, errors.New("not found credential from context")
	}

	if cred.AccessKey == "" {
		return nil, errors.New("accessKey is empty")
	}

	if cred.AccessSecret == "" {
		return nil, errors.New("secretKey is empty")
	}

	if cred.Region == "" {
		return nil, errors.New("region is empty")
	}

	return cred, nil
}
