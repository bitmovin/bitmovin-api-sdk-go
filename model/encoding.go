package model
import (
	"time"
)

type Encoding struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource
	Id string `json:"id,omitempty"`
	CloudRegion CloudRegion `json:"cloudRegion,omitempty"`
	// Version of the encoder
	EncoderVersion string `json:"encoderVersion,omitempty"`
	// Define an external infrastructure to run the encoding on. Note If you set this value, the `cloudRegion` must be 'EXTERNAL'.
	InfrastructureId string `json:"infrastructureId,omitempty"`
	Infrastructure *InfrastructureSettings `json:"infrastructure,omitempty"`
	// You may pass a list of groups associated with this encoding. This will enable you to group results in the statistics resource
	Labels []string `json:"labels,omitempty"`
}

