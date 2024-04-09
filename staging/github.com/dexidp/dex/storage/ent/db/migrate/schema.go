// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthCodesColumns holds the columns for the "auth_codes" table.
	AuthCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "client_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "scopes", Type: field.TypeJSON, Nullable: true},
		{Name: "nonce", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "redirect_uri", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_user_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_username", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_email", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_email_verified", Type: field.TypeBool},
		{Name: "claims_groups", Type: field.TypeJSON, Nullable: true},
		{Name: "claims_preferred_username", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "connector_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "connector_data", Type: field.TypeBytes, Nullable: true},
		{Name: "expiry", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(3)", "postgres": "timestamptz", "sqlite3": "timestamp"}},
		{Name: "code_challenge", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "code_challenge_method", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
	}
	// AuthCodesTable holds the schema information for the "auth_codes" table.
	AuthCodesTable = &schema.Table{
		Name:       "auth_codes",
		Columns:    AuthCodesColumns,
		PrimaryKey: []*schema.Column{AuthCodesColumns[0]},
	}
	// AuthRequestsColumns holds the columns for the "auth_requests" table.
	AuthRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "client_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "scopes", Type: field.TypeJSON, Nullable: true},
		{Name: "response_types", Type: field.TypeJSON, Nullable: true},
		{Name: "redirect_uri", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "nonce", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "state", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "force_approval_prompt", Type: field.TypeBool},
		{Name: "logged_in", Type: field.TypeBool},
		{Name: "claims_user_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_username", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_email", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_email_verified", Type: field.TypeBool},
		{Name: "claims_groups", Type: field.TypeJSON, Nullable: true},
		{Name: "claims_preferred_username", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "connector_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "connector_data", Type: field.TypeBytes, Nullable: true},
		{Name: "expiry", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(3)", "postgres": "timestamptz", "sqlite3": "timestamp"}},
		{Name: "code_challenge", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "code_challenge_method", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "hmac_key", Type: field.TypeBytes},
	}
	// AuthRequestsTable holds the schema information for the "auth_requests" table.
	AuthRequestsTable = &schema.Table{
		Name:       "auth_requests",
		Columns:    AuthRequestsColumns,
		PrimaryKey: []*schema.Column{AuthRequestsColumns[0]},
	}
	// ConnectorsColumns holds the columns for the "connectors" table.
	ConnectorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 100, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "type", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "resource_version", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "config", Type: field.TypeBytes},
	}
	// ConnectorsTable holds the schema information for the "connectors" table.
	ConnectorsTable = &schema.Table{
		Name:       "connectors",
		Columns:    ConnectorsColumns,
		PrimaryKey: []*schema.Column{ConnectorsColumns[0]},
	}
	// DeviceRequestsColumns holds the columns for the "device_requests" table.
	DeviceRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_code", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "device_code", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "client_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "client_secret", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "scopes", Type: field.TypeJSON, Nullable: true},
		{Name: "expiry", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(3)", "postgres": "timestamptz", "sqlite3": "timestamp"}},
	}
	// DeviceRequestsTable holds the schema information for the "device_requests" table.
	DeviceRequestsTable = &schema.Table{
		Name:       "device_requests",
		Columns:    DeviceRequestsColumns,
		PrimaryKey: []*schema.Column{DeviceRequestsColumns[0]},
	}
	// DeviceTokensColumns holds the columns for the "device_tokens" table.
	DeviceTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "device_code", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "status", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "token", Type: field.TypeBytes, Nullable: true},
		{Name: "expiry", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(3)", "postgres": "timestamptz", "sqlite3": "timestamp"}},
		{Name: "last_request", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(3)", "postgres": "timestamptz", "sqlite3": "timestamp"}},
		{Name: "poll_interval", Type: field.TypeInt},
		{Name: "code_challenge", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "code_challenge_method", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
	}
	// DeviceTokensTable holds the schema information for the "device_tokens" table.
	DeviceTokensTable = &schema.Table{
		Name:       "device_tokens",
		Columns:    DeviceTokensColumns,
		PrimaryKey: []*schema.Column{DeviceTokensColumns[0]},
	}
	// KeysColumns holds the columns for the "keys" table.
	KeysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "verification_keys", Type: field.TypeJSON},
		{Name: "signing_key", Type: field.TypeJSON},
		{Name: "signing_key_pub", Type: field.TypeJSON},
		{Name: "next_rotation", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(3)", "postgres": "timestamptz", "sqlite3": "timestamp"}},
	}
	// KeysTable holds the schema information for the "keys" table.
	KeysTable = &schema.Table{
		Name:       "keys",
		Columns:    KeysColumns,
		PrimaryKey: []*schema.Column{KeysColumns[0]},
	}
	// Oauth2clientsColumns holds the columns for the "oauth2clients" table.
	Oauth2clientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 100, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "secret", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "redirect_uris", Type: field.TypeJSON, Nullable: true},
		{Name: "trusted_peers", Type: field.TypeJSON, Nullable: true},
		{Name: "public", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "logo_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
	}
	// Oauth2clientsTable holds the schema information for the "oauth2clients" table.
	Oauth2clientsTable = &schema.Table{
		Name:       "oauth2clients",
		Columns:    Oauth2clientsColumns,
		PrimaryKey: []*schema.Column{Oauth2clientsColumns[0]},
	}
	// OfflineSessionsColumns holds the columns for the "offline_sessions" table.
	OfflineSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "user_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "conn_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "refresh", Type: field.TypeBytes},
		{Name: "connector_data", Type: field.TypeBytes, Nullable: true},
	}
	// OfflineSessionsTable holds the schema information for the "offline_sessions" table.
	OfflineSessionsTable = &schema.Table{
		Name:       "offline_sessions",
		Columns:    OfflineSessionsColumns,
		PrimaryKey: []*schema.Column{OfflineSessionsColumns[0]},
	}
	// PasswordsColumns holds the columns for the "passwords" table.
	PasswordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "hash", Type: field.TypeBytes},
		{Name: "username", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "user_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
	}
	// PasswordsTable holds the schema information for the "passwords" table.
	PasswordsTable = &schema.Table{
		Name:       "passwords",
		Columns:    PasswordsColumns,
		PrimaryKey: []*schema.Column{PasswordsColumns[0]},
	}
	// RefreshTokensColumns holds the columns for the "refresh_tokens" table.
	RefreshTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "client_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "scopes", Type: field.TypeJSON, Nullable: true},
		{Name: "nonce", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_user_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_username", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_email", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "claims_email_verified", Type: field.TypeBool},
		{Name: "claims_groups", Type: field.TypeJSON, Nullable: true},
		{Name: "claims_preferred_username", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "connector_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "connector_data", Type: field.TypeBytes, Nullable: true},
		{Name: "token", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "obsolete_token", Type: field.TypeString, Size: 2147483647, Default: "", SchemaType: map[string]string{"mysql": "varchar(384)", "postgres": "text", "sqlite3": "text"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(3)", "postgres": "timestamptz", "sqlite3": "timestamp"}},
		{Name: "last_used", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(3)", "postgres": "timestamptz", "sqlite3": "timestamp"}},
	}
	// RefreshTokensTable holds the schema information for the "refresh_tokens" table.
	RefreshTokensTable = &schema.Table{
		Name:       "refresh_tokens",
		Columns:    RefreshTokensColumns,
		PrimaryKey: []*schema.Column{RefreshTokensColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthCodesTable,
		AuthRequestsTable,
		ConnectorsTable,
		DeviceRequestsTable,
		DeviceTokensTable,
		KeysTable,
		Oauth2clientsTable,
		OfflineSessionsTable,
		PasswordsTable,
		RefreshTokensTable,
	}
)

func init() {
}