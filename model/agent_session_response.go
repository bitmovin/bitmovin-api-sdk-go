package model

// AgentSessionResponse model
type AgentSessionResponse struct {
	// Session ID (required)
	SessionId *string `json:"sessionId,omitempty"`
	// Agent application name (required)
	AppName *string `json:"appName,omitempty"`
	// User ID (required)
	UserId *string `json:"userId,omitempty"`
}
