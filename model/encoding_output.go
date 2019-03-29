package model

type EncodingOutput struct {
	// Id of the corresponding output
	OutputId string `json:"outputId,omitempty"`
	// Subdirectory where to save the files to
	OutputPath string `json:"outputPath,omitempty"`
	Acl []AclEntry `json:"acl,omitempty"`
}

