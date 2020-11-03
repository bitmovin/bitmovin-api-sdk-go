package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsTsCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsTsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsTsCustomdataAPI constructor for EncodingEncodingsMuxingsTsCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsTsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsTsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTsCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsTsCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTsCustomdataAPI {
    a := &EncodingEncodingsMuxingsTsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get TS muxing Custom Data
func (api *EncodingEncodingsMuxingsTsCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

