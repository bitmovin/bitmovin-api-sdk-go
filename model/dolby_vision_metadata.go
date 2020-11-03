package model


// DolbyVisionMetadata model
type DolbyVisionMetadata struct {
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
    // Dolby Vision Profile (required)
    Profile DolbyVisionProfile `json:"profile,omitempty"`
    // Dolby Vision Metadata Source (required)
    MetadataSource DolbyVisionMetadataSource `json:"metadataSource,omitempty"`
    // ID of the Dolby Vision Metadata Ingest Input Stream which provides the XML Metadata file. Required if metadataSource is set to INPUT_STREAM.
    MetadataInputStreamId *string `json:"metadataInputStreamId,omitempty"`
}



