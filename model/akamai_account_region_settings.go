package model

// AkamaiAccountRegionSettings model
type AkamaiAccountRegionSettings struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Can be freely chosen by the user.
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Id of the VPC subnet for encoding instances (required)
	SubnetId *int64 `json:"subnetId,omitempty"`
	// Id of the firewall for encoding instances (required)
	FirewallId *int64            `json:"firewallId,omitempty"`
	Region     AkamaiCloudRegion `json:"region,omitempty"`
}
