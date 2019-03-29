package model

type AnalysisStartRequest struct {
	Path string `json:"path,omitempty"`
	CloudRegion CloudRegion `json:"cloudRegion,omitempty"`
}

