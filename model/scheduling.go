package model


// Scheduling model
type Scheduling struct {
    // Specify the priority of this encoding (0 - 100). Higher numbers mean higher priority. Default is 50.
    Priority *int32 `json:"priority,omitempty"`
    // List of prewarmed encoder pools. If set, prewarmed encoders from pools with these IDs will be used for the encoding if available. The pool IDs will be tried in the order in which they are passed.
    PrewarmedEncoderPoolIds []string `json:"prewarmedEncoderPoolIds,omitempty"`
}



