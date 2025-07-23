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
	// Deprecated: Amazon access key for legacy support. Use `roleName` instead
	AccessKey *string `json:"accessKey,omitempty"`
	// Deprecated: Amazon secret key for legacy support. Use `roleName` instead
	SecretKey *string `json:"secretKey,omitempty"`
	// Amazon account number (12 digits as per AWS spec) (required)
	AccountNumber *string `json:"accountNumber,omitempty"`
	// Role name including path for the AWS IAM role that will be used by Bitmovin to access the AWS account depicted by `accountNumber`. The role ARN is constructed based on `accountNumber` and `roleName`: `arn:aws:iam::{accountNumber}:role/{roleName}`.  For details on how to create the AWS IAM role in your account, please refer to the [AWS cloud connect setup guide](https://developer.bitmovin.com/encoding/docs/using-bitmovin-cloud-connect-with-aws).
	RoleName *string `json:"roleName,omitempty"`
	// External ID that needs to be set in the trust policy of the AWS IAM role (depicted by `roleName`) that allows Bitmovin access to the AWS account depicted by `accountNumber`
	ExternalId *string `json:"externalId,omitempty"`
}
