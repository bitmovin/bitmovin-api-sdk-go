package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods' endpoints
type EncodingManifestsDashPeriodsAPI struct {
	apiClient *apiclient.APIClient

	// CustomXmlElements communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements' endpoints
	CustomXmlElements *EncodingManifestsDashPeriodsCustomXmlElementsAPI
	// Adaptationsets intermediary API object with no endpoints
	Adaptationsets *EncodingManifestsDashPeriodsAdaptationsetsAPI
}

// NewEncodingManifestsDashPeriodsAPI constructor for EncodingManifestsDashPeriodsAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAPIWithClient constructor for EncodingManifestsDashPeriodsAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAPI {
	a := &EncodingManifestsDashPeriodsAPI{apiClient: apiClient}
	a.CustomXmlElements = NewEncodingManifestsDashPeriodsCustomXmlElementsAPIWithClient(apiClient)
	a.Adaptationsets = NewEncodingManifestsDashPeriodsAdaptationsetsAPIWithClient(apiClient)

	return a
}

// Create Add Period
func (api *EncodingManifestsDashPeriodsAPI) Create(manifestId string, period model.Period) (*model.Period, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.Period
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods", &period, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Period
func (api *EncodingManifestsDashPeriodsAPI) Delete(manifestId string, periodId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Period Details
func (api *EncodingManifestsDashPeriodsAPI) Get(manifestId string, periodId string) (*model.Period, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
	}

	var responseModel model.Period
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Periods
func (api *EncodingManifestsDashPeriodsAPI) List(manifestId string, queryParams ...func(*EncodingManifestsDashPeriodsAPIListQueryParams)) (*pagination.PeriodsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.PeriodsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
