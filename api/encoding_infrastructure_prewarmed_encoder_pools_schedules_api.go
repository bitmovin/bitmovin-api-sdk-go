package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI communicates with '/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}/schedules' endpoints
type EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI constructor for EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI that takes options as argument
func NewEncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPIWithClient constructor for EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI that takes an APIClient as argument
func NewEncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI {
	a := &EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI{apiClient: apiClient}
	return a
}

// Create prewarmed encoder pool schedule
func (api *EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI) Create(poolId string, prewarmedEncoderPoolSchedule model.PrewarmedEncoderPoolSchedule) (*model.PrewarmedEncoderPoolSchedule, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.PrewarmedEncoderPoolSchedule
	err := api.apiClient.Post("/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}/schedules", &prewarmedEncoderPoolSchedule, &responseModel, reqParams)
	return &responseModel, err
}

// Delete prewarmed encoder pool schedule
func (api *EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI) Delete(poolId string, scheduleId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
		params.PathParams["schedule_id"] = scheduleId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}/schedules/{schedule_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Prewarmed encoder pool schedule details
func (api *EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI) Get(poolId string, scheduleId string) (*model.PrewarmedEncoderPoolSchedule, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
		params.PathParams["schedule_id"] = scheduleId
	}

	var responseModel model.PrewarmedEncoderPoolSchedule
	err := api.apiClient.Get("/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}/schedules/{schedule_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List prewarmed encoder pool schedules
func (api *EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI) List(poolId string, queryParams ...func(*EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPIListQueryParams)) (*pagination.PrewarmedEncoderPoolsListPagination, error) {
	queryParameters := &EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.PrewarmedEncoderPoolsListPagination
	err := api.apiClient.Get("/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}/schedules", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
