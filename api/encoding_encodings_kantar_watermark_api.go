package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsKantarWatermarkAPI communicates with '/encoding/encodings/{encoding_id}/kantar-watermark' endpoints
type EncodingEncodingsKantarWatermarkAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsKantarWatermarkAPI constructor for EncodingEncodingsKantarWatermarkAPI that takes options as argument
func NewEncodingEncodingsKantarWatermarkAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsKantarWatermarkAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsKantarWatermarkAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsKantarWatermarkAPIWithClient constructor for EncodingEncodingsKantarWatermarkAPI that takes an APIClient as argument
func NewEncodingEncodingsKantarWatermarkAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsKantarWatermarkAPI {
	a := &EncodingEncodingsKantarWatermarkAPI{apiClient: apiClient}
	return a
}

// Create or replace the Kantar Watermark for an encoding
func (api *EncodingEncodingsKantarWatermarkAPI) Create(encodingId string, kantarWatermark model.KantarWatermark) (*model.KantarWatermark, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.KantarWatermark
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/kantar-watermark", &kantarWatermark, &responseModel, reqParams)
	return &responseModel, err
}

// Delete the Kantar Watermark for an encoding
func (api *EncodingEncodingsKantarWatermarkAPI) Delete(encodingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/kantar-watermark", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get the Kantar Watermark for an encoding
func (api *EncodingEncodingsKantarWatermarkAPI) Get(encodingId string) (*model.KantarWatermark, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.KantarWatermark
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/kantar-watermark", nil, &responseModel, reqParams)
	return &responseModel, err
}
