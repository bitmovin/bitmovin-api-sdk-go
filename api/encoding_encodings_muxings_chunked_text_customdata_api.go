package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsChunkedTextCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/chunked-text/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsChunkedTextCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsChunkedTextCustomdataAPI constructor for EncodingEncodingsMuxingsChunkedTextCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsChunkedTextCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsChunkedTextCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsChunkedTextCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsChunkedTextCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsChunkedTextCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsChunkedTextCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsChunkedTextCustomdataAPI {
	a := &EncodingEncodingsMuxingsChunkedTextCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Chunked Text muxing custom data
func (api *EncodingEncodingsMuxingsChunkedTextCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/chunked-text/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
