package model

import (
	"encoding/json"
)

// S3RoleBasedInput model
type S3RoleBasedInput struct {
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
	// Amazon S3 bucket name (required)
	BucketName *string `json:"bucketName,omitempty"`
	// Amazon ARN of the IAM Role (Identity and Access Management Role) that will be assumed for S3 access.  This role has to be created by the owner of the account with the S3 bucket (i.e., you as a customer). For Bitmovin to be able to assume this role, the following has to be added to the trust policy of the role:  ``` {   \"Effect\": \"Allow\",   \"Principal\": {     \"AWS\": \"arn:aws:iam::630681592166:user/bitmovinCustomerS3Access\"   },   \"Action\": \"sts:AssumeRole\",   \"Condition\": {     \"StringEquals\": {       \"sts:ExternalId\": \"{{externalId}}\"     }   } } ```  where \"arn:aws:iam::630681592166:user/bitmovinCustomerS3Access\" is the Bitmovin user used for the access. The `Condition` is optional but we highly recommend it, see property `externalId` below for more information.  This setup allows Bitmovin assume the provided IAM role and to read data from your S3 bucket. Please note that the IAM role has to have read access on S3.  For more information about role creation please visit https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-service.html#roles-creatingrole-service-console (required)
	RoleArn *string `json:"roleArn,omitempty"`
	// External ID used together with the IAM role identified by `roleArn` to assume S3 access.  This ID is provided by the API if `externalIdMode` is set to `GLOBAL` or `GENERATED`. If present, it has to be added to the trust policy of the IAM role `roleArn` configured above, otherwise the API won't be able to read from the S3 bucket. An appropriate trust policy would look like this:  ``` {   \"Effect\": \"Allow\",   \"Principal\": {     \"AWS\": \"arn:aws:iam::630681592166:user/bitmovinCustomerS3Access\"   },   \"Action\": \"sts:AssumeRole\",   \"Condition\": {     \"StringEquals\": {       \"sts:ExternalId\": \"{{externalId}}\"     }   } } ```  where \"{{externalId}}\" is the generated ID.  This property is optional but we recommend it as an additional security feature. We will use both the `roleArn` and the `externalId` to access your S3 data. If the Amazon IAM role has an external ID configured but it is not provided in the input configuration Bitmovin won't be able to read from the S3 bucket. Also if the external ID does not match the one configured for the IAM role on AWS side, Bitmovin won't be able to access the S3 bucket.  If you need to change the external ID that is used by your IAM role, you need to create a new input, and use the external ID provided by the API to update your IAM role.  For more information please visit https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html
	ExternalId     *string        `json:"externalId,omitempty"`
	ExternalIdMode ExternalIdMode `json:"externalIdMode,omitempty"`
	CloudRegion    AwsCloudRegion `json:"cloudRegion,omitempty"`
}

func (m S3RoleBasedInput) InputType() InputType {
	return InputType_S3_ROLE_BASED
}
func (m S3RoleBasedInput) MarshalJSON() ([]byte, error) {
	type M S3RoleBasedInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "S3_ROLE_BASED"

	return json.Marshal(x)
}
