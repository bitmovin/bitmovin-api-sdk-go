package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingLiveStandbyPoolsAPI communicates with '/encoding/live/standby-pools' endpoints
type EncodingLiveStandbyPoolsAPI struct {
	apiClient *apiclient.APIClient

	// Actions communicates with '/encoding/live/standby-pools/{pool_id}/actions/acquire-encoding' endpoints
	Actions *EncodingLiveStandbyPoolsActionsAPI
	// Encodings communicates with '/encoding/live/standby-pools/{pool_id}/encodings' endpoints
	Encodings *EncodingLiveStandbyPoolsEncodingsAPI
	// Logs communicates with '/encoding/live/standby-pools/{pool_id}/logs' endpoints
	Logs *EncodingLiveStandbyPoolsLogsAPI
}

// NewEncodingLiveStandbyPoolsAPI constructor for EncodingLiveStandbyPoolsAPI that takes options as argument
func NewEncodingLiveStandbyPoolsAPI(options ...apiclient.APIClientOption) (*EncodingLiveStandbyPoolsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveStandbyPoolsAPIWithClient(apiClient), nil
}

// NewEncodingLiveStandbyPoolsAPIWithClient constructor for EncodingLiveStandbyPoolsAPI that takes an APIClient as argument
func NewEncodingLiveStandbyPoolsAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveStandbyPoolsAPI {
	a := &EncodingLiveStandbyPoolsAPI{apiClient: apiClient}
	a.Actions = NewEncodingLiveStandbyPoolsActionsAPIWithClient(apiClient)
	a.Encodings = NewEncodingLiveStandbyPoolsEncodingsAPIWithClient(apiClient)
	a.Logs = NewEncodingLiveStandbyPoolsLogsAPIWithClient(apiClient)

	return a
}

// Create new standby pool
func (api *EncodingLiveStandbyPoolsAPI) Create(liveStandbyPoolRequest model.LiveStandbyPoolRequest) (*model.LiveStandbyPoolDetails, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.LiveStandbyPoolDetails
	err := api.apiClient.Post("/encoding/live/standby-pools", &liveStandbyPoolRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Delete standby pool by id
func (api *EncodingLiveStandbyPoolsAPI) Delete(poolId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/live/standby-pools/{pool_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get details of a standby pool by id
func (api *EncodingLiveStandbyPoolsAPI) Get(poolId string) (*model.LiveStandbyPoolDetails, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.LiveStandbyPoolDetails
	err := api.apiClient.Get("/encoding/live/standby-pools/{pool_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Standby pools
func (api *EncodingLiveStandbyPoolsAPI) List() (*pagination.LiveStandbyPoolResponsesListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel pagination.LiveStandbyPoolResponsesListPagination
	err := api.apiClient.Get("/encoding/live/standby-pools", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Patch Partially update standby pool by id
func (api *EncodingLiveStandbyPoolsAPI) Patch(poolId string, liveStandbyPoolUpdate model.LiveStandbyPoolUpdate) (*model.LiveStandbyPoolDetails, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.LiveStandbyPoolDetails
	err := api.apiClient.Patch("/encoding/live/standby-pools/{pool_id}", &liveStandbyPoolUpdate, &responseModel, reqParams)
	return &responseModel, err
}
