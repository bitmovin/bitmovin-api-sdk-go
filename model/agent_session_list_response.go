package model

// AgentSessionListResponse model
type AgentSessionListResponse struct {
	// Agent application name (required)
	AppName *string `json:"appName,omitempty"`
	// User ID (required)
	UserId *string `json:"userId,omitempty"`
	// Sessions for the user (required)
	Sessions []AgentSessionListItem `json:"sessions,omitempty"`
}
