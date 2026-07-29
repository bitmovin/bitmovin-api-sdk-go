package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AiSceneAnalysisAnalysesAPI communicates with '/ai-scene-analysis/analyses' endpoints
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

// List AI scene analyses
// Returns a paginated list of AI scene analyses.
func (api *AiSceneAnalysisAnalysesAPI) List(queryParams ...func(*AiSceneAnalysisAnalysesAPIListQueryParams)) (*pagination.SceneAnalysisListItemsListPagination, error) {
	queryParameters := &AiSceneAnalysisAnalysesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SceneAnalysisListItemsListPagination
	err := api.apiClient.Get("/ai-scene-analysis/analyses", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AiSceneAnalysisAnalysesAPIListQueryParams contains all query parameters for the List endpoint
type AiSceneAnalysisAnalysesAPIListQueryParams struct {
	Offset        int32                       `query:"offset"`
	Limit         int32                       `query:"limit"`
	Sort          model.SceneAnalysisListSort `query:"sort"`
	CreatedAtFrom model.DateTime              `query:"createdAtFrom"`
	CreatedAtTo   model.DateTime              `query:"createdAtTo"`
}

// Params will return a map of query parameters
func (q *AiSceneAnalysisAnalysesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
