package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsStreamsSpritesCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites/{sprite_id}/customData' endpoints
type EncodingEncodingsStreamsSpritesCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsSpritesCustomdataAPI constructor for EncodingEncodingsStreamsSpritesCustomdataAPI that takes options as argument
func NewEncodingEncodingsStreamsSpritesCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsSpritesCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsSpritesCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsSpritesCustomdataAPIWithClient constructor for EncodingEncodingsStreamsSpritesCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsSpritesCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsSpritesCustomdataAPI {
    a := &EncodingEncodingsStreamsSpritesCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Sprite Custom Data
func (api *EncodingEncodingsStreamsSpritesCustomdataAPI) Get(encodingId string, streamId string, spriteId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["sprite_id"] = spriteId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites/{sprite_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

