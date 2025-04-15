package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AiSceneAnalysisAnalysesAPI intermediary API object with no endpoints
type AiSceneAnalysisAnalysesAPI struct {
	apiClient *apiclient.APIClient

	// ByEncodingId intermediary API object with no endpoints
	ByEncodingId *AiSceneAnalysisAnalysesByEncodingIdAPI
}

// NewAiSceneAnalysisAnalysesAPI constructor for AiSceneAnalysisAnalysesAPI that takes options as argument
func NewAiSceneAnalysisAnalysesAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisAnalysesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisAnalysesAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisAnalysesAPIWithClient constructor for AiSceneAnalysisAnalysesAPI that takes an APIClient as argument
func NewAiSceneAnalysisAnalysesAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisAnalysesAPI {
	a := &AiSceneAnalysisAnalysesAPI{apiClient: apiClient}
	a.ByEncodingId = NewAiSceneAnalysisAnalysesByEncodingIdAPIWithClient(apiClient)

	return a
}
