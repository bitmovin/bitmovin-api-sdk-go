package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingLiveStandbyPoolsEncodingsAPI communicates with '/encoding/live/standby-pools/{pool_id}/encodings' endpoints
type EncodingLiveStandbyPoolsEncodingsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingLiveStandbyPoolsEncodingsAPI constructor for EncodingLiveStandbyPoolsEncodingsAPI that takes options as argument
func NewEncodingLiveStandbyPoolsEncodingsAPI(options ...apiclient.APIClientOption) (*EncodingLiveStandbyPoolsEncodingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveStandbyPoolsEncodingsAPIWithClient(apiClient), nil
}

// NewEncodingLiveStandbyPoolsEncodingsAPIWithClient constructor for EncodingLiveStandbyPoolsEncodingsAPI that takes an APIClient as argument
func NewEncodingLiveStandbyPoolsEncodingsAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveStandbyPoolsEncodingsAPI {
	a := &EncodingLiveStandbyPoolsEncodingsAPI{apiClient: apiClient}
	return a
}

// Delete encoding from pool by id
func (api *EncodingLiveStandbyPoolsEncodingsAPI) Delete(poolId string, id string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
		params.PathParams["id"] = id
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/live/standby-pools/{pool_id}/encodings/{id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List encodings from a standby pool
func (api *EncodingLiveStandbyPoolsEncodingsAPI) List(poolId string, queryParams ...func(*EncodingLiveStandbyPoolsEncodingsAPIListQueryParams)) (*pagination.LiveStandbyPoolEncodingsListPagination, error) {
	queryParameters := &EncodingLiveStandbyPoolsEncodingsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.LiveStandbyPoolEncodingsListPagination
	err := api.apiClient.Get("/encoding/live/standby-pools/{pool_id}/encodings", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingLiveStandbyPoolsEncodingsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingLiveStandbyPoolsEncodingsAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Sort   string `query:"sort"`
	Status string `query:"status"`
}

// Params will return a map of query parameters
func (q *EncodingLiveStandbyPoolsEncodingsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
