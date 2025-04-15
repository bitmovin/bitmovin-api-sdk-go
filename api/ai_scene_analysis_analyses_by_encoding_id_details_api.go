package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AiSceneAnalysisAnalysesByEncodingIdDetailsAPI communicates with '/ai-scene-analysis/analyses/by-encoding-id/{encoding_id}/details' endpoints
type AiSceneAnalysisAnalysesByEncodingIdDetailsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAiSceneAnalysisAnalysesByEncodingIdDetailsAPI constructor for AiSceneAnalysisAnalysesByEncodingIdDetailsAPI that takes options as argument
func NewAiSceneAnalysisAnalysesByEncodingIdDetailsAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisAnalysesByEncodingIdDetailsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisAnalysesByEncodingIdDetailsAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisAnalysesByEncodingIdDetailsAPIWithClient constructor for AiSceneAnalysisAnalysesByEncodingIdDetailsAPI that takes an APIClient as argument
func NewAiSceneAnalysisAnalysesByEncodingIdDetailsAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisAnalysesByEncodingIdDetailsAPI {
	a := &AiSceneAnalysisAnalysesByEncodingIdDetailsAPI{apiClient: apiClient}
	return a
}

// Get AI scene analysis details by encoding ID
// Returns detailed AI scene analysis for a given encoding.
func (api *AiSceneAnalysisAnalysesByEncodingIdDetailsAPI) Get(encodingId string) (*model.SceneAnalysisDetailsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.SceneAnalysisDetailsResponse
	err := api.apiClient.Get("/ai-scene-analysis/analyses/by-encoding-id/{encoding_id}/details", nil, &responseModel, reqParams)
	return &responseModel, err
}
