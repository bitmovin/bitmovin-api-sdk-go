package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AiSceneAnalysisAnalysesByEncodingIdAPI intermediary API object with no endpoints
type AiSceneAnalysisAnalysesByEncodingIdAPI struct {
	apiClient *apiclient.APIClient

	// Details communicates with '/ai-scene-analysis/analyses/by-encoding-id/{encoding_id}/details' endpoints
	Details *AiSceneAnalysisAnalysesByEncodingIdDetailsAPI
}

// NewAiSceneAnalysisAnalysesByEncodingIdAPI constructor for AiSceneAnalysisAnalysesByEncodingIdAPI that takes options as argument
func NewAiSceneAnalysisAnalysesByEncodingIdAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisAnalysesByEncodingIdAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisAnalysesByEncodingIdAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisAnalysesByEncodingIdAPIWithClient constructor for AiSceneAnalysisAnalysesByEncodingIdAPI that takes an APIClient as argument
func NewAiSceneAnalysisAnalysesByEncodingIdAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisAnalysesByEncodingIdAPI {
	a := &AiSceneAnalysisAnalysesByEncodingIdAPI{apiClient: apiClient}
	a.Details = NewAiSceneAnalysisAnalysesByEncodingIdDetailsAPIWithClient(apiClient)

	return a
}
