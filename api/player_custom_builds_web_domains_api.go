package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// PlayerCustomBuildsWebDomainsAPI communicates with '/player/custom-builds/web/domains' endpoints
type PlayerCustomBuildsWebDomainsAPI struct {
    apiClient *apiclient.APIClient
}

// NewPlayerCustomBuildsWebDomainsAPI constructor for PlayerCustomBuildsWebDomainsAPI that takes options as argument
func NewPlayerCustomBuildsWebDomainsAPI(options ...apiclient.APIClientOption) (*PlayerCustomBuildsWebDomainsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerCustomBuildsWebDomainsAPIWithClient(apiClient), nil
}

// NewPlayerCustomBuildsWebDomainsAPIWithClient constructor for PlayerCustomBuildsWebDomainsAPI that takes an APIClient as argument
func NewPlayerCustomBuildsWebDomainsAPIWithClient(apiClient *apiclient.APIClient) *PlayerCustomBuildsWebDomainsAPI {
    a := &PlayerCustomBuildsWebDomainsAPI{apiClient: apiClient}
    return a
}

// Create Add Domain
func (api *PlayerCustomBuildsWebDomainsAPI) Create(customWebPlayerBuildDomain model.CustomWebPlayerBuildDomain) (*model.CustomWebPlayerBuildDomain, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.CustomWebPlayerBuildDomain
    err := api.apiClient.Post("/player/custom-builds/web/domains", &customWebPlayerBuildDomain, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Domain
func (api *PlayerCustomBuildsWebDomainsAPI) Delete(domainId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["domain_id"] = domainId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/player/custom-builds/web/domains/{domain_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Domain Details
func (api *PlayerCustomBuildsWebDomainsAPI) Get(domainId string) (*model.CustomWebPlayerBuildDomain, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["domain_id"] = domainId
    }

    var responseModel model.CustomWebPlayerBuildDomain
    err := api.apiClient.Get("/player/custom-builds/web/domains/{domain_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Domain Details
func (api *PlayerCustomBuildsWebDomainsAPI) List() (*pagination.CustomWebPlayerBuildDomainsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel pagination.CustomWebPlayerBuildDomainsListPagination
    err := api.apiClient.Get("/player/custom-builds/web/domains", nil, &responseModel, reqParams)
    return &responseModel, err
}

