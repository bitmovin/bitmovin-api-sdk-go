package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingLiveStandbyPoolsActionsAPI communicates with '/encoding/live/standby-pools/{pool_id}/actions/acquire-encoding' endpoints
type EncodingLiveStandbyPoolsActionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingLiveStandbyPoolsActionsAPI constructor for EncodingLiveStandbyPoolsActionsAPI that takes options as argument
func NewEncodingLiveStandbyPoolsActionsAPI(options ...apiclient.APIClientOption) (*EncodingLiveStandbyPoolsActionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveStandbyPoolsActionsAPIWithClient(apiClient), nil
}

// NewEncodingLiveStandbyPoolsActionsAPIWithClient constructor for EncodingLiveStandbyPoolsActionsAPI that takes an APIClient as argument
func NewEncodingLiveStandbyPoolsActionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveStandbyPoolsActionsAPI {
	a := &EncodingLiveStandbyPoolsActionsAPI{apiClient: apiClient}
	return a
}

// AcquireEncoding Acquire an encoding from a standby pool
func (api *EncodingLiveStandbyPoolsActionsAPI) AcquireEncoding(poolId string) (*model.LiveStandbyPoolEncoding, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.LiveStandbyPoolEncoding
	err := api.apiClient.Post("/encoding/live/standby-pools/{pool_id}/actions/acquire-encoding", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AcquireEncoding Acquire an encoding from a standby pool
func (api *EncodingLiveStandbyPoolsActionsAPI) AcquireEncodingWithRequestBody(poolId string, liveStandbyPoolAcquireEncoding model.LiveStandbyPoolAcquireEncoding) (*model.LiveStandbyPoolEncoding, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.LiveStandbyPoolEncoding
	err := api.apiClient.Post("/encoding/live/standby-pools/{pool_id}/actions/acquire-encoding", &liveStandbyPoolAcquireEncoding, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteErrorEncodings Delete error encodings from the standby pool
func (api *EncodingLiveStandbyPoolsActionsAPI) DeleteErrorEncodings(poolId string) (*model.LiveStandbyPoolEncoding, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
	}

	var responseModel model.LiveStandbyPoolEncoding
	err := api.apiClient.Post("/encoding/live/standby-pools/{pool_id}/actions/delete-error-encodings", nil, &responseModel, reqParams)
	return &responseModel, err
}
