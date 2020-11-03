package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsInputsAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/inputs' endpoints
type EncodingEncodingsStreamsInputsAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsInputsAPI constructor for EncodingEncodingsStreamsInputsAPI that takes options as argument
func NewEncodingEncodingsStreamsInputsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsInputsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsInputsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsInputsAPIWithClient constructor for EncodingEncodingsStreamsInputsAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsInputsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsInputsAPI {
    a := &EncodingEncodingsStreamsInputsAPI{apiClient: apiClient}
    return a
}

// List Stream Input Analysis Details
func (api *EncodingEncodingsStreamsInputsAPI) List(encodingId string, streamId string) (*pagination.EncodingStreamInputsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel pagination.EncodingStreamInputsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/inputs", nil, &responseModel, reqParams)
    return &responseModel, err
}

