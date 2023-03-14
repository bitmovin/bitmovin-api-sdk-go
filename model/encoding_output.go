package model

// EncodingOutput model
type EncodingOutput struct {
	// Id of the corresponding output (required)
	OutputId *string `json:"outputId,omitempty"`
	// Subdirectory where to save the files to (required)
	OutputPath *string `json:"outputPath,omitempty"`
	// Determines accessibility of files written to this output. Only applies to output types that support ACLs. Defaults to PUBLIC_READ if the list is empty. The destination (e.g. cloud storage bucket) needs to allow the configured ACL
	Acl []AclEntry `json:"acl,omitempty"`
}
