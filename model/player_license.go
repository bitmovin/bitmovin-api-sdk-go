package model
import (
	"time"
)

type PlayerLicense struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	// Name of the resource
	Name string `json:"name,omitempty"`
	// Creation timestamp expressed in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// License Key
	LicenseKey string `json:"licenseKey,omitempty"`
	// Number of impressions recorded
	Impressions *int32 `json:"impressions,omitempty"`
	// Maximum number of impressions
	MaxImpressions *int32 `json:"maxImpressions,omitempty"`
	// Flag if third party licensing is enabled
	ThirdPartyLicensingEnabled *bool `json:"thirdPartyLicensingEnabled,omitempty"`
	// Whitelisted domains
	Domains []Domain `json:"domains,omitempty"`
}

