package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsStreamsCaptionsAPI intermediary API object with no endpoints
type EncodingEncodingsStreamsCaptionsAPI struct {
    apiClient *apiclient.APIClient

    // Cea intermediary API object with no endpoints
    Cea *EncodingEncodingsStreamsCaptionsCeaAPI

}

// NewEncodingEncodingsStreamsCaptionsAPI constructor for EncodingEncodingsStreamsCaptionsAPI that takes options as argument
func NewEncodingEncodingsStreamsCaptionsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsCaptionsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsCaptionsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsCaptionsAPIWithClient constructor for EncodingEncodingsStreamsCaptionsAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsCaptionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsCaptionsAPI {
    a := &EncodingEncodingsStreamsCaptionsAPI{apiClient: apiClient}
    a.Cea = NewEncodingEncodingsStreamsCaptionsCeaAPIWithClient(apiClient)

    return a
}


