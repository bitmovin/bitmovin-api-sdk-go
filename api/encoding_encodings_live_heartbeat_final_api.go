package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveHeartbeatFinalAPI communicates with '/encoding/encodings/{encoding_id}/live/heartbeat-final' endpoints
type EncodingEncodingsLiveHeartbeatFinalAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveHeartbeatFinalAPI constructor for EncodingEncodingsLiveHeartbeatFinalAPI that takes options as argument
func NewEncodingEncodingsLiveHeartbeatFinalAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveHeartbeatFinalAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveHeartbeatFinalAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveHeartbeatFinalAPIWithClient constructor for EncodingEncodingsLiveHeartbeatFinalAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveHeartbeatFinalAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveHeartbeatFinalAPI {
	a := &EncodingEncodingsLiveHeartbeatFinalAPI{apiClient: apiClient}
	return a
}

// Get Final Live Encoding Heartbeat Download URL
// Returns a presigned S3 URL to download the final heartbeat JSON of a live encoding. The URL is valid for 10 minutes. This endpoint is only available once the encoding has reached a final state; calling it while the encoding is still running will result in a 404 response.
func (api *EncodingEncodingsLiveHeartbeatFinalAPI) Get(encodingId string) (*model.LiveEncodingHeartbeatUrlResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.LiveEncodingHeartbeatUrlResponse
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/heartbeat-final", nil, &responseModel, reqParams)
	return &responseModel, err
}
