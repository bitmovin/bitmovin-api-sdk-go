package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id/{id3_tag_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI {
    a := &EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Frame ID ID3 Tag Custom Data of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI) Get(encodingId string, muxingId string, id3TagId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id/{id3_tag_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

