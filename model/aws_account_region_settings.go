package model

// AwsAccountRegionSettings model
type AwsAccountRegionSettings struct {
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
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Limit for the amount of running encodings at a time. Leave empty for no limit.
	LimitParallelEncodings *int64 `json:"limitParallelEncodings,omitempty"`
	// Id of the security group for encoding instances (required)
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
	// Id of the subnet for encoding instances (required)
	SubnetId *string `json:"subnetId,omitempty"`
	// Custom SSH port. Valid values: 1 - 65535. Leave empty if the default SSH port 22 is OK.
	SshPort *int32 `json:"sshPort,omitempty"`
}
