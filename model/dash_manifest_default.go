package model

// DashManifestDefault model
type DashManifestDefault struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
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
	Type       ManifestType            `json:"type,omitempty"`
	// The outputs to store the manifest (required)
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Current status
	Status Status `json:"status,omitempty"`
	// Determines if segmented or progressive representations can appear in the manifest
	Profile DashProfile `json:"profile,omitempty"`
	// The filename of your manifest
	ManifestName *string `json:"manifestName,omitempty"`
	// List of additional XML namespaces to add to the DASH Manifest
	Namespaces []XmlNamespace `json:"namespaces,omitempty"`
	// List of UTC Timings to use for live streaming
	UtcTimings []UtcTiming `json:"utcTimings,omitempty"`
	// The manifest compatibility with the standard DASH Edition.
	DashEditionCompatibility DashEditionCompatibility `json:"dashEditionCompatibility,omitempty"`
	// Determines how timestamps should appear in the manifest
	Iso8601TimestampFormat *DashIso8601TimestampFormat `json:"iso8601TimestampFormat,omitempty"`
	// The minimum buffer time in seconds that the client should maintain to ensure uninterrupted playback. Default is 2 seconds. Note: For VOD ON_DEMAND dash manifests, the default value may differ from 2.0 seconds if not explicitly set.
	MinBufferTime *float64 `json:"minBufferTime,omitempty"`
	// The id of the encoding to create a default manifest for. Either \"encodingId\" or \"periods\" is required.
	EncodingId *string `json:"encodingId,omitempty"`
	// Specifies the algorithm that determines which output of the given encoding is included into the manifest. Note that this is not related to the \"manifestGenerator\" version of the \"Start\" request.
	Version DashManifestDefaultVersion `json:"version,omitempty"`
	// Adds a period for every item. Can only be used when setting \"version\" to \"V2\". Either \"periods\" or \"encodingId\" is required.
	Periods []DefaultDashManifestPeriod `json:"periods,omitempty"`
}
