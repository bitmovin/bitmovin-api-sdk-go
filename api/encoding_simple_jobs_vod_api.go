package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingSimpleJobsVodAPI communicates with '/encoding/simple/jobs/vod' endpoints
type EncodingSimpleJobsVodAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingSimpleJobsVodAPI constructor for EncodingSimpleJobsVodAPI that takes options as argument
func NewEncodingSimpleJobsVodAPI(options ...apiclient.APIClientOption) (*EncodingSimpleJobsVodAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingSimpleJobsVodAPIWithClient(apiClient), nil
}

// NewEncodingSimpleJobsVodAPIWithClient constructor for EncodingSimpleJobsVodAPI that takes an APIClient as argument
func NewEncodingSimpleJobsVodAPIWithClient(apiClient *apiclient.APIClient) *EncodingSimpleJobsVodAPI {
	a := &EncodingSimpleJobsVodAPI{apiClient: apiClient}
	return a
}

// Create a Simple Encoding VOD Job
// Check out our [Simple Encoding API Documentation](https://bitmovin.com/docs/encoding/articles/simple-encoding-api) for additional information about the Simple Encoding API.
func (api *EncodingSimpleJobsVodAPI) Create(simpleEncodingVodJobRequest model.SimpleEncodingVodJobRequest) (*model.SimpleEncodingVodJobResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.SimpleEncodingVodJobResponse
	err := api.apiClient.Post("/encoding/simple/jobs/vod", &simpleEncodingVodJobRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Get Simple Encoding VOD Job details
// Get the details of a Simple VOD Encoding Job.  Check out our [Simple Encoding API Documentation](https://bitmovin.com/docs/encoding/articles/simple-encoding-api) for additional information about the Simple Encoding API.
func (api *EncodingSimpleJobsVodAPI) Get(simpleEncodingJobId string) (*model.SimpleEncodingVodJobResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["simple_encoding_job_id"] = simpleEncodingJobId
	}

	var responseModel model.SimpleEncodingVodJobResponse
	err := api.apiClient.Get("/encoding/simple/jobs/vod/{simple_encoding_job_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Simple Encoding VOD Jobs
func (api *EncodingSimpleJobsVodAPI) List(queryParams ...func(*EncodingSimpleJobsVodAPIListQueryParams)) (*pagination.SimpleEncodingVodJobResponsesListPagination, error) {
	queryParameters := &EncodingSimpleJobsVodAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SimpleEncodingVodJobResponsesListPagination
	err := api.apiClient.Get("/encoding/simple/jobs/vod", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingSimpleJobsVodAPIListQueryParams contains all query parameters for the List endpoint
type EncodingSimpleJobsVodAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Sort   string `query:"sort"`
	Status string `query:"status"`
}

// Params will return a map of query parameters
func (q *EncodingSimpleJobsVodAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
