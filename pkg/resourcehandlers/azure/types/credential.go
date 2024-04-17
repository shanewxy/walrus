package types

import (
	"errors"
)

const (
	SubscriptionID = "subscription_id"
	TenantID       = "tenant_id"
	ClientID       = "client_id"
	ClientSecret   = "client_secret"
)

type Credential struct {
	SubscriptionID string
	TenantID       string
	ClientID       string
	ClientSecret   string
}

func GetCredential(configData map[string][]byte) (*Credential, error) {
	cred := &Credential{}

	cred.SubscriptionID = string(configData[SubscriptionID])
	if cred.SubscriptionID == "" {
		return nil, errors.New("subscriptionID is empty")
	}

	cred.TenantID = string(configData[TenantID])
	if cred.TenantID == "" {
		return nil, errors.New("tenantID is empty")
	}

	cred.ClientID = string(configData[ClientID])
	if cred.ClientID == "" {
		return nil, errors.New("clientID is empty")
	}

	cred.ClientSecret = string(configData[ClientSecret])
	if cred.ClientSecret == "" {
		return nil, errors.New("clientSecret is empty")
	}

	return cred, nil
}
