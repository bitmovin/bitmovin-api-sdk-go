package model

// SceneAnalysisAdPlacementMetadataResponse model
type SceneAnalysisAdPlacementMetadataResponse struct {
	PlacedAds            []AdPosition                         `json:"placedAds,omitempty"`
	AutomaticAdPlacement *AiSceneAnalysisAutomaticAdPlacement `json:"automaticAdPlacement,omitempty"`
}
