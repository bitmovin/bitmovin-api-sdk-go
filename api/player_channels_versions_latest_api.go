package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// PlayerChannelsVersionsLatestAPI communicates with '/player/channels/{channel_name}/versions/latest' endpoints
type PlayerChannelsVersionsLatestAPI struct {
    apiClient *apiclient.APIClient
}

// NewPlayerChannelsVersionsLatestAPI constructor for PlayerChannelsVersionsLatestAPI that takes options as argument
func NewPlayerChannelsVersionsLatestAPI(options ...apiclient.APIClientOption) (*PlayerChannelsVersionsLatestAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerChannelsVersionsLatestAPIWithClient(apiClient), nil
}

// NewPlayerChannelsVersionsLatestAPIWithClient constructor for PlayerChannelsVersionsLatestAPI that takes an APIClient as argument
func NewPlayerChannelsVersionsLatestAPIWithClient(apiClient *apiclient.APIClient) *PlayerChannelsVersionsLatestAPI {
    a := &PlayerChannelsVersionsLatestAPI{apiClient: apiClient}
    return a
}

// Get Latest Player Version for Channel
func (api *PlayerChannelsVersionsLatestAPI) Get(channelName string) (*model.PlayerVersion, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["channel_name"] = channelName
    }

    var responseModel model.PlayerVersion
    err := api.apiClient.Get("/player/channels/{channel_name}/versions/latest", nil, &responseModel, reqParams)
    return &responseModel, err
}

