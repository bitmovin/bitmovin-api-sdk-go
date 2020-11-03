package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsInputStreamsCaptionsAPI intermediary API object with no endpoints
type EncodingEncodingsInputStreamsCaptionsAPI struct {
    apiClient *apiclient.APIClient

    // Cea608 communicates with '/encoding/encodings/{encoding_id}/input-streams/captions/cea608' endpoints
    Cea608 *EncodingEncodingsInputStreamsCaptionsCea608API
    // Cea708 communicates with '/encoding/encodings/{encoding_id}/input-streams/captions/cea708' endpoints
    Cea708 *EncodingEncodingsInputStreamsCaptionsCea708API

}

// NewEncodingEncodingsInputStreamsCaptionsAPI constructor for EncodingEncodingsInputStreamsCaptionsAPI that takes options as argument
func NewEncodingEncodingsInputStreamsCaptionsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsCaptionsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsCaptionsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsCaptionsAPIWithClient constructor for EncodingEncodingsInputStreamsCaptionsAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsCaptionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsCaptionsAPI {
    a := &EncodingEncodingsInputStreamsCaptionsAPI{apiClient: apiClient}
    a.Cea608 = NewEncodingEncodingsInputStreamsCaptionsCea608APIWithClient(apiClient)
    a.Cea708 = NewEncodingEncodingsInputStreamsCaptionsCea708APIWithClient(apiClient)

    return a
}


