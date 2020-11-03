package model


// SmoothStreamingManifest model
type SmoothStreamingManifest struct {
    // Name of the resource. Can be freely chosen by the user.
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
    Type ManifestType `json:"type,omitempty"`
    // The outputs to store the manifest (required)
    Outputs []EncodingOutput `json:"outputs,omitempty"`
    // Current status
    Status Status `json:"status,omitempty"`
    // Filename of the server manifest
    ServerManifestName *string `json:"serverManifestName,omitempty"`
    // Filename of the client manifest
    ClientManifestName *string `json:"clientManifestName,omitempty"`
}



