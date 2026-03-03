package model

import (
	"encoding/json"
)

// AgentChatAttachmentsPart model
type AgentChatAttachmentsPart struct {
	// Attachment payload (required)
	Data *AgentChatAttachmentsData `json:"data,omitempty"`
}

func (m AgentChatAttachmentsPart) AgentChatMessagePartType() AgentChatMessagePartType {
	return AgentChatMessagePartType_DATA_ATTACHMENTS
}
func (m AgentChatAttachmentsPart) MarshalJSON() ([]byte, error) {
	type M AgentChatAttachmentsPart
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "data-attachments"

	return json.Marshal(x)
}
