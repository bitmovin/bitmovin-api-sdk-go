package model

// KantarWatermark model
type KantarWatermark struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Username used to authenticate with the Kantar watermarking service. (required)
	Login *string `json:"login,omitempty"`
	// Password associated with the provided login for authentication. (required)
	Password *string `json:"password,omitempty"`
	// Identifier of the Kantar license to be used when processing the audio watermark. (required)
	LicenseId *int32 `json:"licenseId,omitempty"`
	// Name of the distribution channel associated with the audio content being watermarked. (required)
	ChannelName *string `json:"channelName,omitempty"`
	// Unique reference or identifier for the audio content that will be processed. (required)
	ContentReference *string `json:"contentReference,omitempty"`
	// URL of the Kantar server endpoint used to apply or validate the audio watermark. (required)
	ServerUrl *string `json:"serverUrl,omitempty"`
	// The outputs where the processing report should be delivered. (required)
	ReportOutputs []EncodingOutput `json:"reportOutputs,omitempty"`
}
