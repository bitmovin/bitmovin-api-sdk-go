package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// PlayerChannelsVersionsAPI communicates with '/player/channels/{channel_name}/versions' endpoints
type PlayerChannelsVersionsAPI struct {
    apiClient *apiclient.APIClient

    // Latest communicates with '/player/channels/{channel_name}/versions/latest' endpoints
    Latest *PlayerChannelsVersionsLatestAPI

}

// NewPlayerChannelsVersionsAPI constructor for PlayerChannelsVersionsAPI that takes options as argument
func NewPlayerChannelsVersionsAPI(options ...apiclient.APIClientOption) (*PlayerChannelsVersionsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerChannelsVersionsAPIWithClient(apiClient), nil
}

// NewPlayerChannelsVersionsAPIWithClient constructor for PlayerChannelsVersionsAPI that takes an APIClient as argument
func NewPlayerChannelsVersionsAPIWithClient(apiClient *apiclient.APIClient) *PlayerChannelsVersionsAPI {
    a := &PlayerChannelsVersionsAPI{apiClient: apiClient}
    a.Latest = NewPlayerChannelsVersionsLatestAPIWithClient(apiClient)

    return a
}

// List Player Versions for Channel
func (api *PlayerChannelsVersionsAPI) List(channelName string) (*pagination.PlayerVersionsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["channel_name"] = channelName
    }

    var responseModel pagination.PlayerVersionsListPagination
    err := api.apiClient.Get("/player/channels/{channel_name}/versions", nil, &responseModel, reqParams)
    return &responseModel, err
}

