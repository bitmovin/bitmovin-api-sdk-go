package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// PlayerCustomBuildsWebStatusAPI communicates with '/player/custom-builds/web/{custom_build_id}/status' endpoints
type PlayerCustomBuildsWebStatusAPI struct {
    apiClient *apiclient.APIClient
}

// NewPlayerCustomBuildsWebStatusAPI constructor for PlayerCustomBuildsWebStatusAPI that takes options as argument
func NewPlayerCustomBuildsWebStatusAPI(options ...apiclient.APIClientOption) (*PlayerCustomBuildsWebStatusAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerCustomBuildsWebStatusAPIWithClient(apiClient), nil
}

// NewPlayerCustomBuildsWebStatusAPIWithClient constructor for PlayerCustomBuildsWebStatusAPI that takes an APIClient as argument
func NewPlayerCustomBuildsWebStatusAPIWithClient(apiClient *apiclient.APIClient) *PlayerCustomBuildsWebStatusAPI {
    a := &PlayerCustomBuildsWebStatusAPI{apiClient: apiClient}
    return a
}

// Get Custom Web Player Build Status
func (api *PlayerCustomBuildsWebStatusAPI) Get(customBuildId string) (*model.CustomPlayerBuildStatus, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
    }

    var responseModel model.CustomPlayerBuildStatus
    err := api.apiClient.Get("/player/custom-builds/web/{custom_build_id}/status", nil, &responseModel, reqParams)
    return &responseModel, err
}

