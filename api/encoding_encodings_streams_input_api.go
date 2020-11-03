package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsStreamsInputAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/input' endpoints
type EncodingEncodingsStreamsInputAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsInputAPI constructor for EncodingEncodingsStreamsInputAPI that takes options as argument
func NewEncodingEncodingsStreamsInputAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsInputAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsInputAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsInputAPIWithClient constructor for EncodingEncodingsStreamsInputAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsInputAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsInputAPI {
    a := &EncodingEncodingsStreamsInputAPI{apiClient: apiClient}
    return a
}

// Get Stream Input Details
func (api *EncodingEncodingsStreamsInputAPI) Get(encodingId string, streamId string) (*model.EncodingStreamInputDetails, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel model.EncodingStreamInputDetails
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/input", nil, &responseModel, reqParams)
    return &responseModel, err
}

