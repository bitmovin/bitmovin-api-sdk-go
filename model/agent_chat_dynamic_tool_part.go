package model

import (
	"encoding/json"
)

// AgentChatDynamicToolPart model
type AgentChatDynamicToolPart struct {
	// Tool name (required)
	ToolName *string `json:"toolName,omitempty"`
	// Tool call identifier (required)
	ToolCallId *string `json:"toolCallId,omitempty"`
	// Tool invocation lifecycle state (required)
	State AgentChatDynamicToolState `json:"state,omitempty"`
	// Tool input payload.
	Input *map[string]interface{} `json:"input,omitempty"`
	// Tool output payload.
	Output *map[string]interface{} `json:"output,omitempty"`
	// Error text for failed tool completion.
	ErrorText *string `json:"errorText,omitempty"`
}

func (m AgentChatDynamicToolPart) AgentChatMessagePartType() AgentChatMessagePartType {
	return AgentChatMessagePartType_DYNAMIC_TOOL
}
func (m AgentChatDynamicToolPart) MarshalJSON() ([]byte, error) {
	type M AgentChatDynamicToolPart
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "dynamic-tool"

	return json.Marshal(x)
}
