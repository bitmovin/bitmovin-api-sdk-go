package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AnalyticsLicensesAPI communicates with '/analytics/licenses' endpoints
type AnalyticsLicensesAPI struct {
	apiClient *apiclient.APIClient

	// Domains communicates with '/analytics/licenses/{license_id}/domains' endpoints
	Domains *AnalyticsLicensesDomainsAPI
}

// NewAnalyticsLicensesAPI constructor for AnalyticsLicensesAPI that takes options as argument
func NewAnalyticsLicensesAPI(options ...apiclient.APIClientOption) (*AnalyticsLicensesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsLicensesAPIWithClient(apiClient), nil
}

// NewAnalyticsLicensesAPIWithClient constructor for AnalyticsLicensesAPI that takes an APIClient as argument
func NewAnalyticsLicensesAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsLicensesAPI {
	a := &AnalyticsLicensesAPI{apiClient: apiClient}
	a.Domains = NewAnalyticsLicensesDomainsAPIWithClient(apiClient)

	return a
}

// Create Analytics License
func (api *AnalyticsLicensesAPI) Create(analyticsLicense model.AnalyticsLicense) (*model.AnalyticsLicense, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsLicense
	err := api.apiClient.Post("/analytics/licenses", &analyticsLicense, &responseModel, reqParams)
	return &responseModel, err
}

// Get License
func (api *AnalyticsLicensesAPI) Get(licenseId string) (*model.AnalyticsLicense, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_id"] = licenseId
	}

	var responseModel model.AnalyticsLicense
	err := api.apiClient.Get("/analytics/licenses/{license_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Analytics Licenses
func (api *AnalyticsLicensesAPI) List(queryParams ...func(*AnalyticsLicensesAPIListQueryParams)) (*pagination.AnalyticsLicensesListPagination, error) {
	queryParameters := &AnalyticsLicensesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AnalyticsLicensesListPagination
	err := api.apiClient.Get("/analytics/licenses", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Update Analytics License
func (api *AnalyticsLicensesAPI) Update(licenseId string, analyticsLicense model.AnalyticsLicense) (*model.AnalyticsLicense, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_id"] = licenseId
	}

	var responseModel model.AnalyticsLicense
	err := api.apiClient.Put("/analytics/licenses/{license_id}", &analyticsLicense, &responseModel, reqParams)
	return &responseModel, err
}

// AnalyticsLicensesAPIListQueryParams contains all query parameters for the List endpoint
type AnalyticsLicensesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *AnalyticsLicensesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
