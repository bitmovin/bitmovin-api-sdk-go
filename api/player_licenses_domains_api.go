package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// PlayerLicensesDomainsAPI communicates with '/player/licenses/{license_id}/domains' endpoints
type PlayerLicensesDomainsAPI struct {
    apiClient *apiclient.APIClient
}

// NewPlayerLicensesDomainsAPI constructor for PlayerLicensesDomainsAPI that takes options as argument
func NewPlayerLicensesDomainsAPI(options ...apiclient.APIClientOption) (*PlayerLicensesDomainsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerLicensesDomainsAPIWithClient(apiClient), nil
}

// NewPlayerLicensesDomainsAPIWithClient constructor for PlayerLicensesDomainsAPI that takes an APIClient as argument
func NewPlayerLicensesDomainsAPIWithClient(apiClient *apiclient.APIClient) *PlayerLicensesDomainsAPI {
    a := &PlayerLicensesDomainsAPI{apiClient: apiClient}
    return a
}

// Create Add Domain
func (api *PlayerLicensesDomainsAPI) Create(licenseId string, domain model.Domain) (*model.Domain, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel model.Domain
    err := api.apiClient.Post("/player/licenses/{license_id}/domains", &domain, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Domain
func (api *PlayerLicensesDomainsAPI) Delete(licenseId string, domainId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["license_id"] = licenseId
        params.PathParams["domain_id"] = domainId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/player/licenses/{license_id}/domains/{domain_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List allowed Domains for Player License
func (api *PlayerLicensesDomainsAPI) List(licenseId string) (*pagination.DomainsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel pagination.DomainsListPagination
    err := api.apiClient.Get("/player/licenses/{license_id}/domains", nil, &responseModel, reqParams)
    return &responseModel, err
}

