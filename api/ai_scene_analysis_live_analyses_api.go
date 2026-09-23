package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AiSceneAnalysisLiveAnalysesAPI communicates with '/ai-scene-analysis/live-analyses' endpoints
type AiSceneAnalysisLiveAnalysesAPI struct {
	apiClient *apiclient.APIClient

	// Results intermediary API object with no endpoints
	Results *AiSceneAnalysisLiveAnalysesResultsAPI
}

// NewAiSceneAnalysisLiveAnalysesAPI constructor for AiSceneAnalysisLiveAnalysesAPI that takes options as argument
func NewAiSceneAnalysisLiveAnalysesAPI(options ...apiclient.APIClientOption) (*AiSceneAnalysisLiveAnalysesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAiSceneAnalysisLiveAnalysesAPIWithClient(apiClient), nil
}

// NewAiSceneAnalysisLiveAnalysesAPIWithClient constructor for AiSceneAnalysisLiveAnalysesAPI that takes an APIClient as argument
func NewAiSceneAnalysisLiveAnalysesAPIWithClient(apiClient *apiclient.APIClient) *AiSceneAnalysisLiveAnalysesAPI {
	a := &AiSceneAnalysisLiveAnalysesAPI{apiClient: apiClient}
	a.Results = NewAiSceneAnalysisLiveAnalysesResultsAPIWithClient(apiClient)

	return a
}

// Create Live Analysis
// Creates a Live Analysis. Start the Analysis using the start operation. If creation fails after an Analysis resource has been created, the error response includes a &#x60;Location&#x60; header identifying the Analysis so its failure details can be retrieved.
func (api *AiSceneAnalysisLiveAnalysesAPI) Create(aiSceneAnalysisLiveCreateRequest model.AiSceneAnalysisLiveCreateRequest) (*model.AiSceneAnalysisLiveResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AiSceneAnalysisLiveResponse
	err := api.apiClient.Post("/ai-scene-analysis/live-analyses", &aiSceneAnalysisLiveCreateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Live Analysis
// Deletes a Live Analysis. This is allowed only from &#x60;CREATED&#x60;, &#x60;FINISHED&#x60;, &#x60;CANCELED&#x60;, &#x60;ERROR&#x60;, or &#x60;TRANSFER_ERROR&#x60;. Output resources and files in your storage are not deleted.
func (api *AiSceneAnalysisLiveAnalysesAPI) Delete(analysisId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["analysis_id"] = analysisId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/ai-scene-analysis/live-analyses/{analysis_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Live Analysis details
// Returns the details and current status of a Live Analysis. While the Live Analysis is &#x60;RUNNING&#x60;, the response includes current RTMP ingest details. In all other states, &#x60;ingest&#x60; is omitted.
func (api *AiSceneAnalysisLiveAnalysesAPI) Get(analysisId string) (*model.AiSceneAnalysisLiveResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["analysis_id"] = analysisId
	}

	var responseModel model.AiSceneAnalysisLiveResponse
	err := api.apiClient.Get("/ai-scene-analysis/live-analyses/{analysis_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Live Analyses
// Returns a paginated list of Live Analyses for the effective organization.
func (api *AiSceneAnalysisLiveAnalysesAPI) List(queryParams ...func(*AiSceneAnalysisLiveAnalysesAPIListQueryParams)) (*pagination.AiSceneAnalysisLiveResponsesListPagination, error) {
	queryParameters := &AiSceneAnalysisLiveAnalysesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AiSceneAnalysisLiveResponsesListPagination
	err := api.apiClient.Get("/ai-scene-analysis/live-analyses", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start Live Analysis
// Starts a &#x60;CREATED&#x60; Live Analysis. Repeated calls while it is &#x60;QUEUED&#x60; or &#x60;RUNNING&#x60; reconcile and return the existing resource without launching it again. Poll Get Live Analysis details until the Live Analysis reaches &#x60;RUNNING&#x60;; that response then includes current RTMP ingest details.
func (api *AiSceneAnalysisLiveAnalysesAPI) Start(analysisId string) (*model.AiSceneAnalysisLiveResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["analysis_id"] = analysisId
	}

	var responseModel model.AiSceneAnalysisLiveResponse
	err := api.apiClient.Post("/ai-scene-analysis/live-analyses/{analysis_id}/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Stop Live Analysis
// Requests a stop for a &#x60;QUEUED&#x60; or &#x60;RUNNING&#x60; Live Analysis. Queued cancellation becomes &#x60;CANCELED&#x60;; a graceful running stop becomes &#x60;FINISHED&#x60; only after final required result delivery and usage persistence complete.
func (api *AiSceneAnalysisLiveAnalysesAPI) Stop(analysisId string) (*model.AiSceneAnalysisLiveResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["analysis_id"] = analysisId
	}

	var responseModel model.AiSceneAnalysisLiveResponse
	err := api.apiClient.Post("/ai-scene-analysis/live-analyses/{analysis_id}/stop", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AiSceneAnalysisLiveAnalysesAPIListQueryParams contains all query parameters for the List endpoint
type AiSceneAnalysisLiveAnalysesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *AiSceneAnalysisLiveAnalysesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
