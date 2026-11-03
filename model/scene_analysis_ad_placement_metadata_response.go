package model

// Ad placement results for an encoding.
type SceneAnalysisAdPlacementMetadataResponse struct {
	// Ad breaks selected for this encoding. Empty or omitted when none were selected.
	PlacedAds []AdPosition `json:"placedAds,omitempty"`
	// Placement settings for this encoding. The schedule includes only entries that produced an ad break, with their requested positions.
	AutomaticAdPlacement *AiSceneAnalysisAutomaticAdPlacement `json:"automaticAdPlacement,omitempty"`
}
