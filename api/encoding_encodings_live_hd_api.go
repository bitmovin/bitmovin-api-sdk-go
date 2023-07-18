package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveHdAPI communicates with '/encoding/encodings/{encoding_id}/live/hd/start' endpoints
type EncodingEncodingsLiveHdAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveHdAPI constructor for EncodingEncodingsLiveHdAPI that takes options as argument
func NewEncodingEncodingsLiveHdAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveHdAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveHdAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveHdAPIWithClient constructor for EncodingEncodingsLiveHdAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveHdAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveHdAPI {
	a := &EncodingEncodingsLiveHdAPI{apiClient: apiClient}
	return a
}

// GetStartRequest Live Encoding Start Details
func (api *EncodingEncodingsLiveHdAPI) GetStartRequest(encodingId string) (*model.StartLiveChannelEncodingRequest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.StartLiveChannelEncodingRequest
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/hd/start", nil, &responseModel, reqParams)
	return &responseModel, err
}
