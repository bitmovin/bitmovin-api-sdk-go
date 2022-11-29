package model

import (
	"encoding/json"
)

// S3Output model
type S3Output struct {
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
	// Deprecation notice: This property does not have any effect and will not be returned by GET endpoints
	Acl []AclEntry `json:"acl,omitempty"`
	// Amazon S3 bucket name (required)
	BucketName *string `json:"bucketName,omitempty"`
	// Amazon S3 access key (required)
	AccessKey *string `json:"accessKey,omitempty"`
	// Amazon S3 secret key (required)
	SecretKey *string `json:"secretKey,omitempty"`
	// If set a user defined tag (x-amz-meta-) with that key will be used to store the MD5 hash of the file.
	Md5MetaTag *string `json:"md5MetaTag,omitempty"`
	// The cloud region in which the bucket is located. Is used to determine the ideal location for your encodings automatically.
	CloudRegion AwsCloudRegion `json:"cloudRegion,omitempty"`
	// Specifies the method used for authentication. Must be set to S3_V2 if the region supports both V2 and V4, but the bucket allows V2 only (see https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region)
	SignatureVersion S3SignatureVersion `json:"signatureVersion,omitempty"`
}

func (m S3Output) OutputType() OutputType {
	return OutputType_S3
}
func (m S3Output) MarshalJSON() ([]byte, error) {
	type M S3Output
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "S3"

	return json.Marshal(x)
}
