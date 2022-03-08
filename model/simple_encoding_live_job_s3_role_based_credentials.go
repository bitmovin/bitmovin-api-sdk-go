package model

// SimpleEncodingLiveJobS3RoleBasedCredentials model
type SimpleEncodingLiveJobS3RoleBasedCredentials struct {
	// Amazon ARN of the IAM Role (Identity and Access Management Role) that will be assumed for S3 access. More details can be found [here](https://bitmovin.com/docs/encoding/api-reference/sections/inputs#/Encoding/PostEncodingInputsS3RoleBased) (required)
	RoleArn *string `json:"roleArn,omitempty"`
	// Defines if the organization ID has to be used as externalId when assuming the role. More details can be found [here](https://bitmovin.com/docs/encoding/api-reference/sections/inputs#/Encoding/PostEncodingInputsS3RoleBased)
	UseExternalId *bool `json:"useExternalId,omitempty"`
}
