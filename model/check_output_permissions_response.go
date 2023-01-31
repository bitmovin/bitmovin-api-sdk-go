package model

// CheckOutputPermissionsResponse model
type CheckOutputPermissionsResponse struct {
	// Id of the output for which permissions were checked
	OutputId *string `json:"outputId,omitempty"`
	// The type of the output for which permissions were checked
	OutputType OutputType `json:"outputType,omitempty"`
	// The path on the storage for which permissions were checked. In AWS S3 terminology, this corresponds to a \"prefix\". An empty string corresponds to the root directory.
	Path *string `json:"path,omitempty"`
	// Indicates if permissions on the storage are configured correctly to be used as output target by the Bitmovin encoder. If \"false\", *failureReason* will provide additional information.
	Passed *bool `json:"passed,omitempty"`
	// Contains nothing if the check succeeded. Otherwise, contains a message explaining why it failed.
	FailureReason *string `json:"failureReason,omitempty"`
}
