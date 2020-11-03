package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveTsCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveTsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveTsCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveTsCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveTsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsCustomdataAPI {
    a := &EncodingEncodingsMuxingsProgressiveTsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Progressive TS muxing Custom Data
func (api *EncodingEncodingsMuxingsProgressiveTsCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

