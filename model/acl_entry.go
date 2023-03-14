package model

// AclEntry model
type AclEntry struct {
	// Deprecation notice: The value of this property is not being used. It can be chosen arbitrarily or not set at all
	Scope      *string       `json:"scope,omitempty"`
	Permission AclPermission `json:"permission,omitempty"`
}
