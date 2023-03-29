package model

// StreamsConfigUpdateRequest model
type StreamsConfigUpdateRequest struct {
	// Player style config (required)
	PlayerStyle *map[string]interface{} `json:"playerStyle,omitempty"`
}
