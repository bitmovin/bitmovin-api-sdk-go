package model

import (
	"encoding/json"
)

// DenoiseHqdn3dFilter model
type DenoiseHqdn3dFilter struct {
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
	// A non-negative floating point number which specifies spatial luma strength. It defaults to 4.0.
	LumaSpatial *float64 `json:"lumaSpatial,omitempty"`
	// A non-negative floating point number which specifies spatial chroma strength. It defaults to 3.0*luma_spatial/4.0.
	ChromaSpatial *float64 `json:"chromaSpatial,omitempty"`
	// A floating point number which specifies luma temporal strength. It defaults to 6.0*luma_spatial/4.0.
	LumaTmp *float64 `json:"lumaTmp,omitempty"`
	// A floating point number which specifies chroma temporal strength. It defaults to luma_tmp*chroma_spatial/luma_spatial.
	ChromaTmp *float64 `json:"chromaTmp,omitempty"`
}

func (m DenoiseHqdn3dFilter) FilterType() FilterType {
	return FilterType_DENOISE_HQDN3D
}
func (m DenoiseHqdn3dFilter) MarshalJSON() ([]byte, error) {
	type M DenoiseHqdn3dFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DENOISE_HQDN3D"

	return json.Marshal(x)
}
