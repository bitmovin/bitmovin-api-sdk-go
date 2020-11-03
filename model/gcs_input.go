package model

import (
    "encoding/json"
)

// GcsInput model
type GcsInput struct {
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
    // Name of the bucket (required)
    BucketName *string `json:"bucketName,omitempty"`
    // The cloud region in which the bucket is located. Is used to determine the ideal location for your encodings automatically.
    CloudRegion GoogleCloudRegion `json:"cloudRegion,omitempty"`
    // GCS access key (required)
    AccessKey *string `json:"accessKey,omitempty"`
    // GCS secret key (required)
    SecretKey *string `json:"secretKey,omitempty"`
}

func (m GcsInput) InputType() InputType {
    return InputType_GCS
}
func (m GcsInput) MarshalJSON() ([]byte, error) {
    type M GcsInput
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "GCS"

    return json.Marshal(x)
}


