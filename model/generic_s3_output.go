package model

import (
	"encoding/json"
)

// GenericS3Output model
type GenericS3Output struct {
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
	// Your generic S3 access key (required)
	AccessKey *string `json:"accessKey,omitempty"`
	// Your generic S3 secret key (required)
	SecretKey *string `json:"secretKey,omitempty"`
	// Name of the bucket (required)
	BucketName *string `json:"bucketName,omitempty"`
	// The Generic S3 server hostname (or IP address) (required)
	Host *string `json:"host,omitempty"`
	// The port on which the Generic S3 server is running on (if not provided 8000 will be used)
	Port *int32 `json:"port,omitempty"`
	// Controls whether SSL is used or not
	Ssl *bool `json:"ssl,omitempty"`
	// The signing region to use
	SigningRegion *string `json:"signingRegion,omitempty"`
	// Specifies the method used for authentication
	SignatureVersion S3SignatureVersion `json:"signatureVersion,omitempty"`
	// Specifies the URL access style to use
	AccessStyle S3AccessStyle `json:"accessStyle,omitempty"`
}

func (m GenericS3Output) OutputType() OutputType {
	return OutputType_GENERIC_S3
}
func (m GenericS3Output) MarshalJSON() ([]byte, error) {
	type M GenericS3Output
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "GENERIC_S3"

	return json.Marshal(x)
}
