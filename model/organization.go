package model

// Organization model
type Organization struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Can be freely chosen by the user. (required)
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Specifies the type of the organization in the hierachy. Only sub-organizations can be newly created. (required)
	Type OrganizationType `json:"type,omitempty"`
	// ID of the parent organization
	ParentId *string `json:"parentId,omitempty"`
	// Hexadecimal color
	LabelColor        *string                  `json:"labelColor,omitempty"`
	LimitsPerResource []ResourceLimitContainer `json:"limitsPerResource,omitempty"`
	// which platform initiated organisation creation
	SignupSource SignupSource `json:"signupSource,omitempty"`
	// Flag indicating if MFA is required for the organization
	MfaRequired *bool `json:"mfaRequired,omitempty"`
	// ID of the user who owns the organization
	OwnerUserId *string `json:"ownerUserId,omitempty"`
}
