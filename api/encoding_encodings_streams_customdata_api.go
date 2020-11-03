package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsStreamsCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/customData' endpoints
type EncodingEncodingsStreamsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsCustomdataAPI constructor for EncodingEncodingsStreamsCustomdataAPI that takes options as argument
func NewEncodingEncodingsStreamsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsCustomdataAPIWithClient constructor for EncodingEncodingsStreamsCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsCustomdataAPI {
    a := &EncodingEncodingsStreamsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Stream Custom Data
func (api *EncodingEncodingsStreamsCustomdataAPI) Get(encodingId string, streamId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

