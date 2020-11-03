package model


// PlayerVersion model
type PlayerVersion struct {
    // Id of the resource (required)
    Id *string `json:"id,omitempty"`
    // Version of the Player (required)
    Version *string `json:"version,omitempty"`
    // URL of the specified player (required)
    CdnUrl *string `json:"cdnUrl,omitempty"`
    // Download URL of the specified player package (required)
    DownloadUrl *string `json:"downloadUrl,omitempty"`
    // Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ (required)
    CreatedAt *DateTime `json:"createdAt,omitempty"`
}



