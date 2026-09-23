package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AiSceneAnalysisLiveAnalysesResultsLatestAPI communicates with '/ai-scene-analysis/live-analyses/{analysis_id}/results/latest' endpoints
type AiSceneAnalysisLiveAnalysesResultsLatestAPI struct {
	apiClient *apiclient.APIClient
}

// NewAiSceneAnalysisLiveAnalysesResultsLatestAPI constructor for AiSceneAnalysisLiveAnalysesResultsLatestAPI that takes options as argument
func NewAiSceneAnalysisLiveAnalysesResultsLatestAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisLiveAnalysesResultsLatestAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisLiveAnalysesResultsLatestAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisLiveAnalysesResultsLatestAPIWithClient constructor for AiSceneAnalysisLiveAnalysesResultsLatestAPI that takes an APIClient as argument
func NewAiSceneAnalysisLiveAnalysesResultsLatestAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisLiveAnalysesResultsLatestAPI {
	a := &AiSceneAnalysisLiveAnalysesResultsLatestAPI{apiClient: apiClient}
	return a
}

// Get Live Analysis Latest Result
// Returns the latest cumulative AI analysis results.
func (api *AiSceneAnalysisLiveAnalysesResultsLatestAPI) Get(analysisId string) (*model.AiSceneAnalysisLiveResult, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["analysis_id"] = analysisId
	}

	var responseModel model.AiSceneAnalysisLiveResult
	err := api.apiClient.Get("/ai-scene-analysis/live-analyses/{analysis_id}/results/latest", nil, &responseModel, reqParams)
	return &responseModel, err
}
