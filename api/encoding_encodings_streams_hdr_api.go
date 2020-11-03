package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsStreamsHdrAPI intermediary API object with no endpoints
type EncodingEncodingsStreamsHdrAPI struct {
    apiClient *apiclient.APIClient

    // DolbyVision communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision' endpoints
    DolbyVision *EncodingEncodingsStreamsHdrDolbyVisionAPI

}

// NewEncodingEncodingsStreamsHdrAPI constructor for EncodingEncodingsStreamsHdrAPI that takes options as argument
func NewEncodingEncodingsStreamsHdrAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsHdrAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsHdrAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsHdrAPIWithClient constructor for EncodingEncodingsStreamsHdrAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsHdrAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsHdrAPI {
    a := &EncodingEncodingsStreamsHdrAPI{apiClient: apiClient}
    a.DolbyVision = NewEncodingEncodingsStreamsHdrDolbyVisionAPIWithClient(apiClient)

    return a
}


