package model

// AiSceneAnalysisAssetDescription model
type AiSceneAnalysisAssetDescription struct {
	// Name of the output json file
	Filename *string          `json:"filename,omitempty"`
	Outputs  []EncodingOutput `json:"outputs,omitempty"`
}
