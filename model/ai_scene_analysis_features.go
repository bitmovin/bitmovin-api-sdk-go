package model

// AiSceneAnalysisFeatures model
type AiSceneAnalysisFeatures struct {
	// AI scene analysis will create an asset description file.
	AssetDescription *AiSceneAnalysisAssetDescription `json:"assetDescription,omitempty"`
	// Ad markers placed on detected scene changes and configured positions.
	AutomaticAdPlacement *AiSceneAnalysisAutomaticAdPlacement `json:"automaticAdPlacement,omitempty"`
}
