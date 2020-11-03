package model


// EgressInformation model
type EgressInformation struct {
    Category EgressCategory `json:"category,omitempty"`
    // The number of bytes that have been transferred to the output (required)
    Bytes *int64 `json:"bytes,omitempty"`
    OutputType OutputType `json:"outputType,omitempty"`
    OutputRegion CloudRegion `json:"outputRegion,omitempty"`
}



