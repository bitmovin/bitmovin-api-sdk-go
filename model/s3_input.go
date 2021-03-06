package model

import (
	"encoding/json"
)

// S3Input model
type S3Input struct {
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
	// The cloud region in which the bucket is located. Is used to determine the ideal location for your encodings automatically.
	CloudRegion AwsCloudRegion `json:"cloudRegion,omitempty"`
	// Name of the bucket (required)
	BucketName *string `json:"bucketName,omitempty"`
	// Amazon access key (required)
	AccessKey *string `json:"accessKey,omitempty"`
	// Amazon secret key (required)
	SecretKey *string `json:"secretKey,omitempty"`
}

func (m S3Input) InputType() InputType {
	return InputType_S3
}
func (m S3Input) MarshalJSON() ([]byte, error) {
	type M S3Input
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "S3"

	return json.Marshal(x)
}
