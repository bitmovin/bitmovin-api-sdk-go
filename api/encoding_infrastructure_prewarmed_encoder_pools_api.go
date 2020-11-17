package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructurePrewarmedEncoderPoolsAPI communicates with '/encoding/infrastructure/prewarmed-encoder-pools' endpoints
type EncodingInfrastructurePrewarmedEncoderPoolsAPI struct {
	apiClient *apiclient.APIClient

	// Schedules communicates with '/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}/schedules' endpoints
	Schedules *EncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPI
}

// NewEncodingInfrastructurePrewarmedEncoderPoolsAPI constructor for EncodingInfrastructurePrewarmedEncoderPoolsAPI that takes options as argument
func NewEncodingInfrastructurePrewarmedEncoderPoolsAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructurePrewarmedEncoderPoolsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructurePrewarmedEncoderPoolsAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructurePrewarmedEncoderPoolsAPIWithClient constructor for EncodingInfrastructurePrewarmedEncoderPoolsAPI that takes an APIClient as argument
func NewEncodingInfrastructurePrewarmedEncoderPoolsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructurePrewarmedEncoderPoolsAPI {
	a := &EncodingInfrastructurePrewarmedEncoderPoolsAPI{apiClient: apiClient}
	a.Schedules = NewEncodingInfrastructurePrewarmedEncoderPoolsSchedulesAPIWithClient(apiClient)

	return a
}

// Create prewarmed encoder pool
func (api *EncodingInfrastructurePrewarmedEncoderPoolsAPI) Create(prewarmedEncoderPool model.PrewarmedEncoderPool) (*model.PrewarmedEncoderPool, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.PrewarmedEncoderPool
	err := api.apiClient.Post("/encoding/infrastructure/prewarmed-encoder-pools", &prewarmedEncoderPool, &responseModel, reqParams)
	return &responseModel, err
}

// Delete prewarmed encoder pool
func (api *EncodingInfrastructurePrewarmedEncoderPoolsAPI) Delete(poolId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Prewarmed encoder pool details
func (api *EncodingInfrastructurePrewarmedEncoderPoolsAPI) Get(poolId string) (*model.PrewarmedEncoderPool, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.PrewarmedEncoderPool
	err := api.apiClient.Get("/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List prewarmed encoder pools
func (api *EncodingInfrastructurePrewarmedEncoderPoolsAPI) List(queryParams ...func(*EncodingInfrastructurePrewarmedEncoderPoolsAPIListQueryParams)) (*pagination.PrewarmedEncoderPoolsListPagination, error) {
	queryParameters := &EncodingInfrastructurePrewarmedEncoderPoolsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.PrewarmedEncoderPoolsListPagination
	err := api.apiClient.Get("/encoding/infrastructure/prewarmed-encoder-pools", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start prewarmed encoder pool
func (api *EncodingInfrastructurePrewarmedEncoderPoolsAPI) Start(poolId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Stop prewarmed encoder pool
func (api *EncodingInfrastructurePrewarmedEncoderPoolsAPI) Stop(poolId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/infrastructure/prewarmed-encoder-pools/{pool_id}/stop", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructurePrewarmedEncoderPoolsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructurePrewarmedEncoderPoolsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructurePrewarmedEncoderPoolsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
