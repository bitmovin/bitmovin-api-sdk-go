package model

// Available language variants of an analysis result.
type SceneAnalysisLanguagesResponse struct {
	// Language codes registered for available analysis details. The base English result can be included even when no translations were requested; this is not the list of detected input languages.
	OutputLanguageCodes []string `json:"outputLanguageCodes,omitempty"`
}
