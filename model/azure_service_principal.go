package model

// Azure service principal credentials for Microsoft Entra ID authentication. This authentication method is usable from encoder version 2.273.0 onwards.
type AzureServicePrincipal struct {
	// Tenant ID (Directory ID) of the Azure service principal (required)
	TenantId *string `json:"tenantId,omitempty"`
	// Client ID of the Azure service principal (required)
	ClientId *string `json:"clientId,omitempty"`
	// Client secret of the Azure service principal
	ClientSecret *string `json:"clientSecret,omitempty"`
	// PEM-encoded client certificate and private key of the Azure service principal. Newline symbols must be preserved.
	ClientCertificate *string `json:"clientCertificate,omitempty"`
}
