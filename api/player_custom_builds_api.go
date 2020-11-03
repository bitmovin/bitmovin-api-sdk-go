package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// PlayerCustomBuildsAPI intermediary API object with no endpoints
type PlayerCustomBuildsAPI struct {
    apiClient *apiclient.APIClient

    // Web communicates with '/player/custom-builds/web' endpoints
    Web *PlayerCustomBuildsWebAPI

}

// NewPlayerCustomBuildsAPI constructor for PlayerCustomBuildsAPI that takes options as argument
func NewPlayerCustomBuildsAPI(options ...apiclient.APIClientOption) (*PlayerCustomBuildsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerCustomBuildsAPIWithClient(apiClient), nil
}

// NewPlayerCustomBuildsAPIWithClient constructor for PlayerCustomBuildsAPI that takes an APIClient as argument
func NewPlayerCustomBuildsAPIWithClient(apiClient *apiclient.APIClient) *PlayerCustomBuildsAPI {
    a := &PlayerCustomBuildsAPI{apiClient: apiClient}
    a.Web = NewPlayerCustomBuildsWebAPIWithClient(apiClient)

    return a
}


