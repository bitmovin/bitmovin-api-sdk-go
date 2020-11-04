package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMxfCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mxf/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsMxfCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMxfCustomdataAPI constructor for EncodingEncodingsMuxingsMxfCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsMxfCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMxfCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMxfCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMxfCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsMxfCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMxfCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMxfCustomdataAPI {
	a := &EncodingEncodingsMuxingsMxfCustomdataAPI{apiClient: apiClient}
	return a
}

// Get MXF muxing Custom Data
func (api *EncodingEncodingsMuxingsMxfCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mxf/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
