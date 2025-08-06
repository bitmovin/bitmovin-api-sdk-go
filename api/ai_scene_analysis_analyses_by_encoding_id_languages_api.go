package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AiSceneAnalysisAnalysesByEncodingIdLanguagesAPI communicates with '/ai-scene-analysis/analyses/by-encoding-id/{encoding_id}/languages' endpoints
type AiSceneAnalysisAnalysesByEncodingIdLanguagesAPI struct {
	apiClient *apiclient.APIClient
}

// NewAiSceneAnalysisAnalysesByEncodingIdLanguagesAPI constructor for AiSceneAnalysisAnalysesByEncodingIdLanguagesAPI that takes options as argument
func NewAiSceneAnalysisAnalysesByEncodingIdLanguagesAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisAnalysesByEncodingIdLanguagesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisAnalysesByEncodingIdLanguagesAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisAnalysesByEncodingIdLanguagesAPIWithClient constructor for AiSceneAnalysisAnalysesByEncodingIdLanguagesAPI that takes an APIClient as argument
func NewAiSceneAnalysisAnalysesByEncodingIdLanguagesAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisAnalysesByEncodingIdLanguagesAPI {
	a := &AiSceneAnalysisAnalysesByEncodingIdLanguagesAPI{apiClient: apiClient}
	return a
}

// Get AI scene analysis languages by encoding ID
// Returns list of languages from AI scene analysis for a given encoding.
func (api *AiSceneAnalysisAnalysesByEncodingIdLanguagesAPI) Get(encodingId string) (*model.SceneAnalysisLanguagesResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.SceneAnalysisLanguagesResponse
	err := api.apiClient.Get("/ai-scene-analysis/analyses/by-encoding-id/{encoding_id}/languages", nil, &responseModel, reqParams)
	return &responseModel, err
}
