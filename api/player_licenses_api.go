package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// PlayerLicensesAPI communicates with '/player/licenses' endpoints
type PlayerLicensesAPI struct {
    apiClient *apiclient.APIClient

    // Analytics communicates with '/player/licenses/{license_id}/analytics' endpoints
    Analytics *PlayerLicensesAnalyticsAPI
    // Domains communicates with '/player/licenses/{license_id}/domains' endpoints
    Domains *PlayerLicensesDomainsAPI
    // ThirdPartyLicensing communicates with '/player/licenses/{license_id}/third-party-licensing' endpoints
    ThirdPartyLicensing *PlayerLicensesThirdPartyLicensingAPI

}

// NewPlayerLicensesAPI constructor for PlayerLicensesAPI that takes options as argument
func NewPlayerLicensesAPI(options ...apiclient.APIClientOption) (*PlayerLicensesAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewPlayerLicensesAPIWithClient(apiClient), nil
}

// NewPlayerLicensesAPIWithClient constructor for PlayerLicensesAPI that takes an APIClient as argument
func NewPlayerLicensesAPIWithClient(apiClient *apiclient.APIClient) *PlayerLicensesAPI {
    a := &PlayerLicensesAPI{apiClient: apiClient}
    a.Analytics = NewPlayerLicensesAnalyticsAPIWithClient(apiClient)
    a.Domains = NewPlayerLicensesDomainsAPIWithClient(apiClient)
    a.ThirdPartyLicensing = NewPlayerLicensesThirdPartyLicensingAPIWithClient(apiClient)

    return a
}

// Create Player License
func (api *PlayerLicensesAPI) Create(playerLicense model.PlayerLicense) (*model.PlayerLicense, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.PlayerLicense
    err := api.apiClient.Post("/player/licenses", &playerLicense, &responseModel, reqParams)
    return &responseModel, err
}
// Get License
func (api *PlayerLicensesAPI) Get(licenseId string) (*model.PlayerLicense, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel model.PlayerLicense
    err := api.apiClient.Get("/player/licenses/{license_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Player Licenses
func (api *PlayerLicensesAPI) List(queryParams ...func(*PlayerLicensesAPIListQueryParams)) (*pagination.PlayerLicensesListPagination, error) {
    queryParameters := &PlayerLicensesAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.PlayerLicensesListPagination
    err := api.apiClient.Get("/player/licenses", nil, &responseModel, reqParams)
    return &responseModel, err
}

// PlayerLicensesAPIListQueryParams contains all query parameters for the List endpoint
type PlayerLicensesAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *PlayerLicensesAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


