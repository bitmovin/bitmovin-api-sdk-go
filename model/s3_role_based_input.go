package model
import (
	"time"
)

type S3RoleBasedInput struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp expressed in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp expressed in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource
	Id string `json:"id,omitempty"`
	// Amazon S3 bucket name
	BucketName string `json:"bucketName,omitempty"`
	// Amazon ARN of the Role that will be assumed for S3 access.
	RoleArn string `json:"roleArn,omitempty"`
	CloudRegion AwsCloudRegion `json:"cloudRegion,omitempty"`
}
func (o S3RoleBasedInput) InputType() InputType {
    return InputType_S3_ROLE_BASED
}

