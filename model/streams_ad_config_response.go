package model

// StreamsAdConfigResponse model
type StreamsAdConfigResponse struct {
	// The identifier of the streams ad config
	Id  *string             `json:"id,omitempty"`
	Ads []StreamsAdConfigAd `json:"ads,omitempty"`
}
