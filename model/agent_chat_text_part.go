package model

import (
	"bytes"
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

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
