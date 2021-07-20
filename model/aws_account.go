package model

// AwsAccount model
type AwsAccount struct {
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
	// Amazon access key (required)
	AccessKey *string `json:"accessKey,omitempty"`
	// Amazon secret key (required)
	SecretKey *string `json:"secretKey,omitempty"`
	// Amazon account number (12 digits as per AWS spec) (required)
	AccountNumber *string `json:"accountNumber,omitempty"`
}
