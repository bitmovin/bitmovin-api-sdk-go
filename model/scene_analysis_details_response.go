package model

// SceneAnalysisDetailsResponse model
type SceneAnalysisDetailsResponse struct {
	Scenes                      []Scene   `json:"scenes,omitempty"`
	Duration                    *float32  `json:"duration,omitempty"`
	Description                 *string   `json:"description,omitempty"`
	Keywords                    []string  `json:"keywords,omitempty"`
	Ratings                     []Rating  `json:"ratings,omitempty"`
	SensitiveTopics             []string  `json:"sensitiveTopics,omitempty"`
	IabSensitiveTopicTaxonomies []string  `json:"iabSensitiveTopicTaxonomies,omitempty"`
	InputLanguageCodes          []string  `json:"inputLanguageCodes,omitempty"`
	Metadata                    *Metadata `json:"metadata,omitempty"`
}
