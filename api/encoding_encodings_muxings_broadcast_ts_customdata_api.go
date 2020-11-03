package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsBroadcastTsCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsBroadcastTsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsBroadcastTsCustomdataAPI constructor for EncodingEncodingsMuxingsBroadcastTsCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsBroadcastTsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsBroadcastTsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsBroadcastTsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsBroadcastTsCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsBroadcastTsCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsBroadcastTsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsBroadcastTsCustomdataAPI {
    a := &EncodingEncodingsMuxingsBroadcastTsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Broadcast TS muxing Custom Data
func (api *EncodingEncodingsMuxingsBroadcastTsCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

