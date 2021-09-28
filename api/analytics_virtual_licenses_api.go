package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AnalyticsVirtualLicensesAPI communicates with '/analytics/virtual-licenses' endpoints
type AnalyticsVirtualLicensesAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsVirtualLicensesAPI constructor for AnalyticsVirtualLicensesAPI that takes options as argument
func NewAnalyticsVirtualLicensesAPI(options ...apiclient.APIClientOption) (*AnalyticsVirtualLicensesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsVirtualLicensesAPIWithClient(apiClient), nil
}

// NewAnalyticsVirtualLicensesAPIWithClient constructor for AnalyticsVirtualLicensesAPI that takes an APIClient as argument
func NewAnalyticsVirtualLicensesAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsVirtualLicensesAPI {
	a := &AnalyticsVirtualLicensesAPI{apiClient: apiClient}
	return a
}

// Create Analytics Virtual License
func (api *AnalyticsVirtualLicensesAPI) Create(analyticsVirtualLicenseRequest model.AnalyticsVirtualLicenseRequest) (*model.AnalyticsVirtualLicense, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsVirtualLicense
	err := api.apiClient.Post("/analytics/virtual-licenses", &analyticsVirtualLicenseRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Analytics Virtual License
func (api *AnalyticsVirtualLicensesAPI) Delete(virtualLicenseId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["virtual_license_id"] = virtualLicenseId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/analytics/virtual-licenses/{virtual_license_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Analytics Virtual Licenses
func (api *AnalyticsVirtualLicensesAPI) List(queryParams ...func(*AnalyticsVirtualLicensesAPIListQueryParams)) (*pagination.AnalyticsVirtualLicensesListPagination, error) {
	queryParameters := &AnalyticsVirtualLicensesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AnalyticsVirtualLicensesListPagination
	err := api.apiClient.Get("/analytics/virtual-licenses", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Update Analytics Virtual License
func (api *AnalyticsVirtualLicensesAPI) Update(virtualLicenseId string, analyticsVirtualLicenseRequest model.AnalyticsVirtualLicenseRequest) (*model.AnalyticsVirtualLicense, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["virtual_license_id"] = virtualLicenseId
	}

	var responseModel model.AnalyticsVirtualLicense
	err := api.apiClient.Put("/analytics/virtual-licenses/{virtual_license_id}", &analyticsVirtualLicenseRequest, &responseModel, reqParams)
	return &responseModel, err
}

// AnalyticsVirtualLicensesAPIListQueryParams contains all query parameters for the List endpoint
type AnalyticsVirtualLicensesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *AnalyticsVirtualLicensesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
