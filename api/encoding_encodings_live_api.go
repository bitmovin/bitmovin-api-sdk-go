package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveAPI communicates with '/encoding/encodings/{encoding_id}/live' endpoints
type EncodingEncodingsLiveAPI struct {
	apiClient *apiclient.APIClient

	// InsertableContent communicates with '/encoding/encodings/{encoding_id}/live/insertable-content' endpoints
	InsertableContent *EncodingEncodingsLiveInsertableContentAPI
}

// NewEncodingEncodingsLiveAPI constructor for EncodingEncodingsLiveAPI that takes options as argument
func NewEncodingEncodingsLiveAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveAPIWithClient constructor for EncodingEncodingsLiveAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveAPI {
	a := &EncodingEncodingsLiveAPI{apiClient: apiClient}
	a.InsertableContent = NewEncodingEncodingsLiveInsertableContentAPIWithClient(apiClient)

	return a
}

// Get Live Encoding Details
func (api *EncodingEncodingsLiveAPI) Get(encodingId string) (*model.LiveEncoding, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.LiveEncoding
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetStartRequest Live Encoding Start Details
func (api *EncodingEncodingsLiveAPI) GetStartRequest(encodingId string) (*model.StartLiveEncodingRequest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.StartLiveEncodingRequest
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Restart Re-Start Live Encoding
func (api *EncodingEncodingsLiveAPI) Restart(encodingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/restart", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start Live Encoding
func (api *EncodingEncodingsLiveAPI) Start(encodingId string, startLiveEncodingRequest model.StartLiveEncodingRequest) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/start", &startLiveEncodingRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Stop Live Encoding
func (api *EncodingEncodingsLiveAPI) Stop(encodingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/stop", nil, &responseModel, reqParams)
	return &responseModel, err
}