package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AnalyticsAlertingIncidentsAPI communicates with '/analytics/alerting/incidents' endpoints
type AnalyticsAlertingIncidentsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsAlertingIncidentsAPI constructor for AnalyticsAlertingIncidentsAPI that takes options as argument
func NewAnalyticsAlertingIncidentsAPI(options ...apiclient.APIClientOption) (*AnalyticsAlertingIncidentsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsAlertingIncidentsAPIWithClient(apiClient), nil
}

// NewAnalyticsAlertingIncidentsAPIWithClient constructor for AnalyticsAlertingIncidentsAPI that takes an APIClient as argument
func NewAnalyticsAlertingIncidentsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAlertingIncidentsAPI {
	a := &AnalyticsAlertingIncidentsAPI{apiClient: apiClient}
	return a
}

// List Incidents
func (api *AnalyticsAlertingIncidentsAPI) List(queryParams ...func(*AnalyticsAlertingIncidentsAPIListQueryParams)) (*pagination.AnalyticsIncidentsListPagination, error) {
	queryParameters := &AnalyticsAlertingIncidentsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AnalyticsIncidentsListPagination
	err := api.apiClient.Get("/analytics/alerting/incidents", nil, &responseModel, reqParams)
	return &responseModel, err
}

// ListByLicenseKey List Incidents per license
func (api *AnalyticsAlertingIncidentsAPI) ListByLicenseKey(licenseKey string, queryParams ...func(*AnalyticsAlertingIncidentsAPIListByLicenseKeyQueryParams)) (*pagination.AnalyticsIncidentsListByLicenseKeyPagination, error) {
	queryParameters := &AnalyticsAlertingIncidentsAPIListByLicenseKeyQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_key"] = licenseKey
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AnalyticsIncidentsListByLicenseKeyPagination
	err := api.apiClient.Get("/analytics/alerting/incidents/{license_key}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AnalyticsAlertingIncidentsAPIListQueryParams contains all query parameters for the List endpoint
type AnalyticsAlertingIncidentsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *AnalyticsAlertingIncidentsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}

// AnalyticsAlertingIncidentsAPIListByLicenseKeyQueryParams contains all query parameters for the ListByLicenseKey endpoint
type AnalyticsAlertingIncidentsAPIListByLicenseKeyQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *AnalyticsAlertingIncidentsAPIListByLicenseKeyQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
