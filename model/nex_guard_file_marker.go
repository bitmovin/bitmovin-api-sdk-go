package model

// NexGuardFileMarker model
type NexGuardFileMarker struct {
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
	// Use the base64 license string that Nagra provides you. (required)
	License *string `json:"license,omitempty"`
	// The type of watermarking to be used: * `OTT` - A/B watermarking (for video streams only) * `DUPLICATED` - Stream duplication to match A/B video streams in CDN delivery (for audio streams only)
	WatermarkType NexGuardWatermarkingType `json:"watermarkType,omitempty"`
	// Specify the payload ID that you want to be associated with this output. Valid values vary depending on your Nagra NexGuard forensic watermarking workflow. For PreRelease Content (NGPR), specify an integer from 1 through 4,194,303. You must generate a unique ID for each asset you watermark, and keep a record of th ID. Neither Nagra nor Bitmovin keep track of this for you.
	Payload *int32 `json:"payload,omitempty"`
	// Enter one of the watermarking preset strings that Nagra provides you.
	Preset *string `json:"preset,omitempty"`
	// Optional. Ignore this setting unless Nagra support directs you to specify a value. When you don't specify a value here, the Nagra NexGuard library uses its default value.
	Strength NexGuardWatermarkingStrength `json:"strength,omitempty"`
}
