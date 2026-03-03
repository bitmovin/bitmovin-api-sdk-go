package model

// AgentChatAttachment model
type AgentChatAttachment struct {
	// Attachment kind (required)
	Kind *string `json:"kind,omitempty"`
	// Tool call identifier
	ToolCallId *string `json:"toolCallId,omitempty"`
	// Attachment attributes map (required)
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
}
