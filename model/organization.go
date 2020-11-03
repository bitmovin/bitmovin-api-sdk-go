package model


// Organization model
type Organization struct {
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
    // Id of the resource (required)
    Id *string `json:"id,omitempty"`
    Type OrganizationType `json:"type,omitempty"`
    // ID of the parent organization
    ParentId *string `json:"parentId,omitempty"`
    // Hexadecimal color
    LabelColor *string `json:"labelColor,omitempty"`
    LimitsPerResource []ResourceLimitContainer `json:"limitsPerResource,omitempty"`
}



