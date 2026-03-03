package model

import (
	"encoding/json"
)

// AgentChatTextPart model
type AgentChatTextPart struct {
	// Text message content (required)
	Text *string `json:"text,omitempty"`
}

func (m AgentChatTextPart) AgentChatMessagePartType() AgentChatMessagePartType {
	return AgentChatMessagePartType_TEXT
}
func (m AgentChatTextPart) MarshalJSON() ([]byte, error) {
	type M AgentChatTextPart
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "text"

	return json.Marshal(x)
}
