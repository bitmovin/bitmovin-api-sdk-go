package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AiSceneAnalysisAPI intermediary API object with no endpoints
type AiSceneAnalysisAPI struct {
	apiClient *apiclient.APIClient

	// Analyses intermediary API object with no endpoints
	Analyses *AiSceneAnalysisAnalysesAPI
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

	return a
}
