package model

// SimpleEncodingVodJobResponse model
type SimpleEncodingVodJobResponse struct {
	// The identifier of the Simple Encoding VOD Job
	Id *string `json:"id,omitempty"`
	// The current status of the Simple Encoding VOD Job
	Status SimpleEncodingVodJobStatus `json:"status,omitempty"`
	// The template that has been used for the encoding.
	EncodingTemplate EncodingTemplate `json:"encodingTemplate,omitempty"`
	// The identifier of the encoding that has been created based on the job request. This is only returned once the job execution has been successful and the encoding could be started.
	EncodingId *string                         `json:"encodingId,omitempty"`
	Inputs     []SimpleEncodingVodJobUrlInput  `json:"inputs,omitempty"`
	Outputs    []SimpleEncodingVodJobUrlOutput `json:"outputs,omitempty"`
	// Describes all the errors in cases the status of the job is 'error'.
	Errors []SimpleEncodingVodJobErrors `json:"errors,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ. The job is updated for example when the status changes
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// This property will be used for naming the encoding and the manifests.
	Name *string `json:"name,omitempty"`
}
