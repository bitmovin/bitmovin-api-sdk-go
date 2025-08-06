package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI communicates with '/ai-scene-analysis/analyses/by-encoding-id/{encoding_id}/details/language/{language_code}' endpoints
type AiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI struct {
	apiClient *apiclient.APIClient
}

// NewAiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI constructor for AiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI that takes options as argument
func NewAiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPIWithClient constructor for AiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI that takes an APIClient as argument
func NewAiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI {
	a := &AiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI{apiClient: apiClient}
	return a
}

// Get translated AI scene analysis details by encoding ID and language code
// Returns detailed translated AI scene analysis for a given encoding.
func (api *AiSceneAnalysisAnalysesByEncodingIdDetailsLanguageAPI) Get(encodingId string, languageCode string) (*model.SceneAnalysisDetailsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["language_code"] = languageCode
	}

	var responseModel model.SceneAnalysisDetailsResponse
	err := api.apiClient.Get("/ai-scene-analysis/analyses/by-encoding-id/{encoding_id}/details/language/{language_code}", nil, &responseModel, reqParams)
	return &responseModel, err
}
