package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsStreamsThumbnailsCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}/customData' endpoints
type EncodingEncodingsStreamsThumbnailsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsThumbnailsCustomdataAPI constructor for EncodingEncodingsStreamsThumbnailsCustomdataAPI that takes options as argument
func NewEncodingEncodingsStreamsThumbnailsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsThumbnailsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsThumbnailsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsThumbnailsCustomdataAPIWithClient constructor for EncodingEncodingsStreamsThumbnailsCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsThumbnailsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsThumbnailsCustomdataAPI {
    a := &EncodingEncodingsStreamsThumbnailsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Thumbnail Custom Data
func (api *EncodingEncodingsStreamsThumbnailsCustomdataAPI) Get(encodingId string, streamId string, thumbnailId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["thumbnail_id"] = thumbnailId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

