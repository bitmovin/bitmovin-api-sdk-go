package model

// AzureSpeechToCaptionsSettings model
type AzureSpeechToCaptionsSettings struct {
	// Credential settings to access Azure Speech Services
	AzureSpeechServicesCredentials *AzureSpeechServicesCredentials `json:"azureSpeechServicesCredentials,omitempty"`
	// Azure Speech Services Region Identifier. The list of speech service supported regions can be found at Azure's official documentation.
	Region *string `json:"region,omitempty"`
	// Azure Speech Services API endpoint. This information can be found in Azure's Speech resource data.
	ApiEndpoint *string `json:"apiEndpoint,omitempty"`
	// Azure Speech to captions supported language (IETF BCP 47 language tag). The list of supported languages can be found at Azure's official documentation.
	Language *string `json:"language,omitempty"`
	// How many MILLISECONDS to delay the display of each caption, to mimic a real-time experience. The minimum value is 0.
	CaptionDelay *int64 `json:"captionDelay,omitempty"`
	// How many MILLISECONDS a caption should remain on screen if it is not replaced by another. The minimum value is 0.
	CaptionRemainTime *int64 `json:"captionRemainTime,omitempty"`
	// The maximum number of characters per line for a caption.  The minimum value is 20.
	CaptionMaxLineLength *int64 `json:"captionMaxLineLength,omitempty"`
	// The number of lines for a caption. The minimum value is 1.
	CaptionLines *int64 `json:"captionLines,omitempty"`
	// The profanity filter options are:  - Masked: Replaces letters in profane words with asterisk (*) characters. - Raw: Include the profane words verbatim. - Removed: Removes profane words.
	ProfanityOption AzureSpeechToCaptionsProfanity `json:"profanityOption,omitempty"`
}
