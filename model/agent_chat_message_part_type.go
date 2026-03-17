package model

// AgentChatMessagePartType : AgentChatMessagePartType model
type AgentChatMessagePartType string

// List of possible AgentChatMessagePartType values
const (
	AgentChatMessagePartType_TEXT             AgentChatMessagePartType = "text"
	AgentChatMessagePartType_DATA_ATTACHMENTS AgentChatMessagePartType = "data-attachments"
	AgentChatMessagePartType_DYNAMIC_TOOL     AgentChatMessagePartType = "dynamic-tool"
)
