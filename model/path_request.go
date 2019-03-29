package model

type PathRequest struct {
	Path string `json:"path,omitempty"`
	Recursive *bool `json:"recursive,omitempty"`
	CloudRegion CloudRegion `json:"cloudRegion,omitempty"`
}

