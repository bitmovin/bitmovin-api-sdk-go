package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// AgentChatMessage model
type AgentChatMessage struct {
	// Message ID (required)
	Id *string `json:"id,omitempty"`
	// Message role (required)
	Role *string `json:"role,omitempty"`
	// Message parts (required)
	Parts []AgentChatMessagePart `json:"parts,omitempty"`
}

// UnmarshalJSON unmarshals model AgentChatMessage from a JSON structure
func (m *AgentChatMessage) UnmarshalJSON(raw []byte) error {
	var data struct {
		Id    *string         `json:"id"`
		Role  *string         `json:"role"`
		Parts json.RawMessage `json:"parts"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result AgentChatMessage

	result.Id = data.Id
	result.Role = data.Role

	allOfParts, err := UnmarshalAgentChatMessagePartSlice(bytes.NewBuffer(data.Parts), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Parts = allOfParts

	*m = result

	return nil
}
