package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AiSceneAnalysisAPI intermediary API object with no endpoints
type AiSceneAnalysisAPI struct {
	apiClient *apiclient.APIClient

	// Analyses communicates with '/ai-scene-analysis/analyses' endpoints
	Analyses *AiSceneAnalysisAnalysesAPI
	// LiveAnalyses communicates with '/ai-scene-analysis/live-analyses' endpoints
	LiveAnalyses *AiSceneAnalysisLiveAnalysesAPI
}

// NewAiSceneAnalysisAPI constructor for AiSceneAnalysisAPI that takes options as argument
func NewAiSceneAnalysisAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisAPIWithClient constructor for AiSceneAnalysisAPI that takes an APIClient as argument
func NewAiSceneAnalysisAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisAPI {
	a := &AiSceneAnalysisAPI{apiClient: apiClient}
	a.Analyses = NewAiSceneAnalysisAnalysesAPIWithClient(apiClient)
	a.LiveAnalyses = NewAiSceneAnalysisLiveAnalysesAPIWithClient(apiClient)

	return a
}
