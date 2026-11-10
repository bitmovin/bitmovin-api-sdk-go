package model

// AiSceneAnalysisAssetDescription model
type AiSceneAnalysisAssetDescription struct {
	// Name of the output JSON file. The [asset-description JSON Schema](https://storage.googleapis.com/ai-scene-analysis-demo-content/schemas/asset-description.schema.json) describes the structure of newly written original and translated files.
	Filename *string `json:"filename,omitempty"`
	// Output destinations in which the generated asset-description JSON files are stored.
	Outputs []EncodingOutput `json:"outputs,omitempty"`
}
