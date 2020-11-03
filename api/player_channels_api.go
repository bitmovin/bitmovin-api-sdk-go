package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// PlayerChannelsAPI communicates with '/player/channels' endpoints
type PlayerChannelsAPI struct {
    apiClient *apiclient.APIClient

    // Versions communicates with '/player/channels/{channel_name}/versions' endpoints
    Versions *PlayerChannelsVersionsAPI

}

// NewPlayerChannelsAPI constructor for PlayerChannelsAPI that takes options as argument
func NewPlayerChannelsAPI(options ...apiclient.APIClientOption) (*PlayerChannelsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerChannelsAPIWithClient(apiClient), nil
}

// NewPlayerChannelsAPIWithClient constructor for PlayerChannelsAPI that takes an APIClient as argument
func NewPlayerChannelsAPIWithClient(apiClient *apiclient.APIClient) *PlayerChannelsAPI {
    a := &PlayerChannelsAPI{apiClient: apiClient}
    a.Versions = NewPlayerChannelsVersionsAPIWithClient(apiClient)

    return a
}

// List Player Channels
func (api *PlayerChannelsAPI) List() (*pagination.PlayerChannelsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel pagination.PlayerChannelsListPagination
    err := api.apiClient.Get("/player/channels", nil, &responseModel, reqParams)
    return &responseModel, err
}

