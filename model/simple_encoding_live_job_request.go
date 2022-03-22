package model

// SimpleEncodingLiveJobRequest model
type SimpleEncodingLiveJobRequest struct {
	Input *SimpleEncodingLiveJobInput `json:"input,omitempty"`
	// output of the live encoding job (required)
	Outputs []SimpleEncodingLiveJobUrlOutput `json:"outputs,omitempty"`
	// The cloud region that will be used for the live encoding. This value has to be set.
	CloudRegion SimpleEncodingLiveCloudRegion `json:"cloudRegion,omitempty"`
	// This property will be used for naming the encoding.
	Name *string `json:"name,omitempty"`
}
