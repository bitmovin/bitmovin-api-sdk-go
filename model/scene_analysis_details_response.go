package model

// SceneAnalysisDetailsResponse model
type SceneAnalysisDetailsResponse struct {
	Scenes                      []Scene  `json:"scenes,omitempty"`
	Description                 *string  `json:"description,omitempty"`
	Keywords                    []string `json:"keywords,omitempty"`
	Ratings                     []Rating `json:"ratings,omitempty"`
	SensitiveTopics             []string `json:"sensitiveTopics,omitempty"`
	IabSensitiveTopicTaxonomies []string `json:"iabSensitiveTopicTaxonomies,omitempty"`
}
