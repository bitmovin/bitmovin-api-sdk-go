package model

// AgentSessionHistoryResponse model
type AgentSessionHistoryResponse struct {
	// Session ID (required)
	SessionId *string `json:"sessionId,omitempty"`
	// Agent application name (required)
	AppName *string `json:"appName,omitempty"`
	// User ID (required)
	UserId *string `json:"userId,omitempty"`
	// Session title
	Title *string `json:"title,omitempty"`
	// Session message history (required)
	Messages []AgentChatMessage `json:"messages,omitempty"`
}
