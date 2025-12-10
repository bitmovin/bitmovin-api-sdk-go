package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveHeartbeatAPI communicates with '/encoding/encodings/{encoding_id}/live/heartbeat' endpoints
type EncodingEncodingsLiveHeartbeatAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveHeartbeatAPI constructor for EncodingEncodingsLiveHeartbeatAPI that takes options as argument
func NewEncodingEncodingsLiveHeartbeatAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveHeartbeatAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveHeartbeatAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveHeartbeatAPIWithClient constructor for EncodingEncodingsLiveHeartbeatAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveHeartbeatAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveHeartbeatAPI {
	a := &EncodingEncodingsLiveHeartbeatAPI{apiClient: apiClient}
	return a
}

// Get Live Encoding Heartbeat
func (api *EncodingEncodingsLiveHeartbeatAPI) Get(encodingId string) (*model.LiveEncodingHeartbeat, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.LiveEncodingHeartbeat
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/heartbeat", nil, &responseModel, reqParams)
	return &responseModel, err
}
