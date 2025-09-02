package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI communicates with '/ai-scene-analysis/analyses/by-encoding-id/{encoding_id}/ad-placements' endpoints
type AiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI constructor for AiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI that takes options as argument
func NewAiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPIWithClient constructor for AiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI that takes an APIClient as argument
func NewAiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI {
	a := &AiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI{apiClient: apiClient}
	return a
}

// Get AI scene analysis ad placements by encoding ID
// Returns ad placements from AI scene analysis for a given encoding.
func (api *AiSceneAnalysisAnalysesByEncodingIdAdPlacementsAPI) Get(encodingId string) (*model.SceneAnalysisAdPlacementMetadataResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.SceneAnalysisAdPlacementMetadataResponse
	err := api.apiClient.Get("/ai-scene-analysis/analyses/by-encoding-id/{encoding_id}/ad-placements", nil, &responseModel, reqParams)
	return &responseModel, err
}
