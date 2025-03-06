package model

// AiContentAnalysisFeatures model
type AiContentAnalysisFeatures struct {
	// AI content analysis will create an asset description file.
	AssetDescription *AiContentAnalysisAssetDescription `json:"assetDescription,omitempty"`
	// Ad markers placed on detected scene changes and configured positions.
	AutomaticAdPlacement *AiContentAnalysisAutomaticAdPlacement `json:"automaticAdPlacement,omitempty"`
}
