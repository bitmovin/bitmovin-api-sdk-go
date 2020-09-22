package model

// ResourceType : ResourceType model
type ResourceType string

// List of possible ResourceType values
const (
	ResourceType_ACCOUNT   ResourceType = "account"
	ResourceType_ENCODING  ResourceType = "encoding"
	ResourceType_PLAYER    ResourceType = "player"
	ResourceType_ANALYTICS ResourceType = "analytics"
)
