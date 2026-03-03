package model

// AgentChatAttachmentsData model
type AgentChatAttachmentsData struct {
	// Attachment list (required)
	Attachments []AgentChatAttachment `json:"attachments,omitempty"`
}
