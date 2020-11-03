package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// PlayerCustomBuildsWebAPI communicates with '/player/custom-builds/web' endpoints
type PlayerCustomBuildsWebAPI struct {
    apiClient *apiclient.APIClient

    // Domains communicates with '/player/custom-builds/web/domains' endpoints
    Domains *PlayerCustomBuildsWebDomainsAPI
    // Status communicates with '/player/custom-builds/web/{custom_build_id}/status' endpoints
    Status *PlayerCustomBuildsWebStatusAPI
    // Download communicates with '/player/custom-builds/web/{custom_build_id}/download' endpoints
    Download *PlayerCustomBuildsWebDownloadAPI

}

// NewPlayerCustomBuildsWebAPI constructor for PlayerCustomBuildsWebAPI that takes options as argument
func NewPlayerCustomBuildsWebAPI(options ...apiclient.APIClientOption) (*PlayerCustomBuildsWebAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerCustomBuildsWebAPIWithClient(apiClient), nil
}

// NewPlayerCustomBuildsWebAPIWithClient constructor for PlayerCustomBuildsWebAPI that takes an APIClient as argument
func NewPlayerCustomBuildsWebAPIWithClient(apiClient *apiclient.APIClient) *PlayerCustomBuildsWebAPI {
    a := &PlayerCustomBuildsWebAPI{apiClient: apiClient}
    a.Domains = NewPlayerCustomBuildsWebDomainsAPIWithClient(apiClient)
    a.Status = NewPlayerCustomBuildsWebStatusAPIWithClient(apiClient)
    a.Download = NewPlayerCustomBuildsWebDownloadAPIWithClient(apiClient)

    return a
}

// Create Add Custom Web Player Build
func (api *PlayerCustomBuildsWebAPI) Create(customPlayerBuildDetails model.CustomPlayerBuildDetails) (*model.CustomPlayerBuildDetails, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.CustomPlayerBuildDetails
    err := api.apiClient.Post("/player/custom-builds/web", &customPlayerBuildDetails, &responseModel, reqParams)
    return &responseModel, err
}
// Get Custom Web Player Build Details
func (api *PlayerCustomBuildsWebAPI) Get(customBuildId string) (*model.CustomPlayerBuildStatus, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
    }

    var responseModel model.CustomPlayerBuildStatus
    err := api.apiClient.Get("/player/custom-builds/web/{custom_build_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Custom Web Player Builds
func (api *PlayerCustomBuildsWebAPI) List() (*pagination.CustomPlayerBuildDetailssListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel pagination.CustomPlayerBuildDetailssListPagination
    err := api.apiClient.Get("/player/custom-builds/web", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Start Custom Web Player Build
func (api *PlayerCustomBuildsWebAPI) Start(customBuildId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Post("/player/custom-builds/web/{custom_build_id}/start", nil, &responseModel, reqParams)
    return &responseModel, err
}

