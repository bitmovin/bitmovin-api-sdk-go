package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingLiveStreamKeysActionsAPI communicates with '/encoding/live/stream-keys/actions/unassign' endpoints
type EncodingLiveStreamKeysActionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingLiveStreamKeysActionsAPI constructor for EncodingLiveStreamKeysActionsAPI that takes options as argument
func NewEncodingLiveStreamKeysActionsAPI(options ...apiclient.APIClientOption) (*EncodingLiveStreamKeysActionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveStreamKeysActionsAPIWithClient(apiClient), nil
}

// NewEncodingLiveStreamKeysActionsAPIWithClient constructor for EncodingLiveStreamKeysActionsAPI that takes an APIClient as argument
func NewEncodingLiveStreamKeysActionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveStreamKeysActionsAPI {
	a := &EncodingLiveStreamKeysActionsAPI{apiClient: apiClient}
	return a
}

// Unassign stream keys
func (api *EncodingLiveStreamKeysActionsAPI) Unassign(streamKeysUnassignAction model.StreamKeysUnassignAction) (*model.StreamKeysUnassignAction, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.StreamKeysUnassignAction
	err := api.apiClient.Post("/encoding/live/stream-keys/actions/unassign", &streamKeysUnassignAction, &responseModel, reqParams)
	return &responseModel, err
}
