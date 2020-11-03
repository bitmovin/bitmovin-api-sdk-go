package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// PlayerCustomBuildsWebDownloadAPI communicates with '/player/custom-builds/web/{custom_build_id}/download' endpoints
type PlayerCustomBuildsWebDownloadAPI struct {
    apiClient *apiclient.APIClient
}

// NewPlayerCustomBuildsWebDownloadAPI constructor for PlayerCustomBuildsWebDownloadAPI that takes options as argument
func NewPlayerCustomBuildsWebDownloadAPI(options ...apiclient.APIClientOption) (*PlayerCustomBuildsWebDownloadAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerCustomBuildsWebDownloadAPIWithClient(apiClient), nil
}

// NewPlayerCustomBuildsWebDownloadAPIWithClient constructor for PlayerCustomBuildsWebDownloadAPI that takes an APIClient as argument
func NewPlayerCustomBuildsWebDownloadAPIWithClient(apiClient *apiclient.APIClient) *PlayerCustomBuildsWebDownloadAPI {
    a := &PlayerCustomBuildsWebDownloadAPI{apiClient: apiClient}
    return a
}

// Get Custom Web Player Build Download Link
func (api *PlayerCustomBuildsWebDownloadAPI) Get(customBuildId string) (*model.CustomPlayerBuildDownload, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
    }

    var responseModel model.CustomPlayerBuildDownload
    err := api.apiClient.Get("/player/custom-builds/web/{custom_build_id}/download", nil, &responseModel, reqParams)
    return &responseModel, err
}

