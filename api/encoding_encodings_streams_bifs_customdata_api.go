package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsStreamsBifsCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs/{bif_id}/customData' endpoints
type EncodingEncodingsStreamsBifsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsBifsCustomdataAPI constructor for EncodingEncodingsStreamsBifsCustomdataAPI that takes options as argument
func NewEncodingEncodingsStreamsBifsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsBifsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsBifsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsBifsCustomdataAPIWithClient constructor for EncodingEncodingsStreamsBifsCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsBifsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsBifsCustomdataAPI {
    a := &EncodingEncodingsStreamsBifsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Bif Custom Data
func (api *EncodingEncodingsStreamsBifsCustomdataAPI) Get(encodingId string, streamId string, bifId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["bif_id"] = bifId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs/{bif_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

