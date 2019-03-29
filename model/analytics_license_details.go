package model
import (
	"time"
)

type AnalyticsLicenseDetails struct {
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
	// License Key
	LicenseKey string `json:"licenseKey,omitempty"`
	// Maximum number of datapoints
	MaxDatapoints *int64 `json:"maxDatapoints,omitempty"`
	// Number of datapoints recorded
	Datapoints *int64 `json:"datapoints,omitempty"`
	// Whitelisted domains
	Domains []AnalyticsLicenseDomain `json:"domains,omitempty"`
}

