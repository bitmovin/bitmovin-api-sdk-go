package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AiSceneAnalysisLiveAnalysesResultsAPI intermediary API object with no endpoints
type AiSceneAnalysisLiveAnalysesResultsAPI struct {
	apiClient *apiclient.APIClient

	// Latest communicates with '/ai-scene-analysis/live-analyses/{analysis_id}/results/latest' endpoints
	Latest *AiSceneAnalysisLiveAnalysesResultsLatestAPI
}

// NewAiSceneAnalysisLiveAnalysesResultsAPI constructor for AiSceneAnalysisLiveAnalysesResultsAPI that takes options as argument
func NewAiSceneAnalysisLiveAnalysesResultsAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisLiveAnalysesResultsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisLiveAnalysesResultsAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisLiveAnalysesResultsAPIWithClient constructor for AiSceneAnalysisLiveAnalysesResultsAPI that takes an APIClient as argument
func NewAiSceneAnalysisLiveAnalysesResultsAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisLiveAnalysesResultsAPI {
	a := &AiSceneAnalysisLiveAnalysesResultsAPI{apiClient: apiClient}
	a.Latest = NewAiSceneAnalysisLiveAnalysesResultsLatestAPIWithClient(apiClient)

	return a
}
