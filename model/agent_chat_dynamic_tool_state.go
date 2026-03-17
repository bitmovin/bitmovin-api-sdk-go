package model

// AgentChatDynamicToolState : AgentChatDynamicToolState model
type AgentChatDynamicToolState string

// List of possible AgentChatDynamicToolState values
const (
	AgentChatDynamicToolState_INPUT_STREAMING  AgentChatDynamicToolState = "input-streaming"
	AgentChatDynamicToolState_INPUT_AVAILABLE  AgentChatDynamicToolState = "input-available"
	AgentChatDynamicToolState_OUTPUT_AVAILABLE AgentChatDynamicToolState = "output-available"
	AgentChatDynamicToolState_OUTPUT_ERROR     AgentChatDynamicToolState = "output-error"
)
