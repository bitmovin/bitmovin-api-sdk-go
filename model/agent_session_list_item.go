package model

// AgentSessionListItem model
type AgentSessionListItem struct {
	// Session ID (required)
	SessionId *string `json:"sessionId,omitempty"`
	// Session title
	Title *string `json:"title,omitempty"`
	// Last update time in seconds
	LastUpdateTimeSeconds *float64 `json:"lastUpdateTimeSeconds,omitempty"`
}
