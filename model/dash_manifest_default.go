package model


// V2 of the default dash manifest is an experimental feature and might be subject to change in the future. 
type DashManifestDefault struct {
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
    Profile DashProfile `json:"profile,omitempty"`
    // The filename of your manifest
    ManifestName *string `json:"manifestName,omitempty"`
    // List of additional XML namespaces to add to the DASH Manifest
    Namespaces []XmlNamespace `json:"namespaces,omitempty"`
    // List of UTC Timings to use for live streaming
    UtcTimings []UtcTiming `json:"utcTimings,omitempty"`
    // The id of the encoding to create a default manifest from. Required: encodingId or periods
    EncodingId *string `json:"encodingId,omitempty"`
    // The version of the default manifest generator
    Version DashManifestDefaultVersion `json:"version,omitempty"`
    // Adds a period for every item. Required: encodingId or periods
    Periods []DefaultDashManifestPeriod `json:"periods,omitempty"`
}



