package model

// AiSceneAnalysisAssetDescription model
type AiSceneAnalysisAssetDescription struct {
	// Name of the output json file
	Filename *string `json:"filename,omitempty"`
	// Output destinations in which the generated asset-description JSON files are stored.
	Outputs []EncodingOutput `json:"outputs,omitempty"`
}
