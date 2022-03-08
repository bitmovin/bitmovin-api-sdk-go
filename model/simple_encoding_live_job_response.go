package model

// SimpleEncodingLiveJobResponse model
type SimpleEncodingLiveJobResponse struct {
	// The identifier of the Simple Encoding Live Job
	Id *string `json:"id,omitempty"`
	// The current status of the Simple Encoding Live Job
	Status SimpleEncodingLiveJobStatus `json:"status,omitempty"`
	// The identifier of the encoding that has been created based on the job request. This is only returned once the job execution has been successful and the encoding could be started.
	EncodingId *string `json:"encodingId,omitempty"`
	// The IP address of the encoder for this job request. This is only returned once the job execution has been successful and the encoding could be started. Ingest is expected to be sent to this IP address.
	EncoderIp *string `json:"encoderIp,omitempty"`
	// Stream key of the live encoder
	StreamKey *string                          `json:"streamKey,omitempty"`
	Input     *SimpleEncodingLiveJobInput      `json:"input,omitempty"`
	Outputs   []SimpleEncodingLiveJobUrlOutput `json:"outputs,omitempty"`
	// Describes all the errors in cases the status of the job is 'error'.   TODO right now this is the same for VOD and LIVE? maybe rename the schema and use the same for both?
	Errors []SimpleEncodingVodJobErrors `json:"errors,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ. The job is updated for example when the status changes
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// This property will be used for naming the encoding and the manifests.
	Name *string `json:"name,omitempty"`
	// The cloud region that will be used for the live encoding
	CloudRegion SimpleEncodingLiveCloudRegion `json:"cloudRegion,omitempty"`
}
