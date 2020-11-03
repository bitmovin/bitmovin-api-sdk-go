package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp3CustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsMp3CustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp3CustomdataAPI constructor for EncodingEncodingsMuxingsMp3CustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp3CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp3CustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp3CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp3CustomdataAPIWithClient constructor for EncodingEncodingsMuxingsMp3CustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp3CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp3CustomdataAPI {
    a := &EncodingEncodingsMuxingsMp3CustomdataAPI{apiClient: apiClient}
    return a
}

// Get MP3 muxing Custom Data
func (api *EncodingEncodingsMuxingsMp3CustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

