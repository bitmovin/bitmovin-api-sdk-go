package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsLicensesDomainsAPI communicates with '/analytics/licenses/{license_id}/domains' endpoints
type AnalyticsLicensesDomainsAPI struct {
    apiClient *apiclient.APIClient
}

// NewAnalyticsLicensesDomainsAPI constructor for AnalyticsLicensesDomainsAPI that takes options as argument
func NewAnalyticsLicensesDomainsAPI(options ...apiclient.APIClientOption) (*AnalyticsLicensesDomainsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsLicensesDomainsAPIWithClient(apiClient), nil
}

// NewAnalyticsLicensesDomainsAPIWithClient constructor for AnalyticsLicensesDomainsAPI that takes an APIClient as argument
func NewAnalyticsLicensesDomainsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsLicensesDomainsAPI {
    a := &AnalyticsLicensesDomainsAPI{apiClient: apiClient}
    return a
}

// Create Add Domain
func (api *AnalyticsLicensesDomainsAPI) Create(licenseId string, domain model.Domain) (*model.Domain, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel model.Domain
    err := api.apiClient.Post("/analytics/licenses/{license_id}/domains", &domain, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Domain
func (api *AnalyticsLicensesDomainsAPI) Delete(licenseId string, domainId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["license_id"] = licenseId
        params.PathParams["domain_id"] = domainId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/analytics/licenses/{license_id}/domains/{domain_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get List License Domains
func (api *AnalyticsLicensesDomainsAPI) Get(licenseId string) (*model.DomainList, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel model.DomainList
    err := api.apiClient.Get("/analytics/licenses/{license_id}/domains", nil, &responseModel, reqParams)
    return &responseModel, err
}

