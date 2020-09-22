package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsCustomXmlElementsAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements' endpoints
type EncodingManifestsDashPeriodsCustomXmlElementsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsCustomXmlElementsAPI constructor for EncodingManifestsDashPeriodsCustomXmlElementsAPI that takes options as argument
func NewEncodingManifestsDashPeriodsCustomXmlElementsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsCustomXmlElementsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsCustomXmlElementsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsCustomXmlElementsAPIWithClient constructor for EncodingManifestsDashPeriodsCustomXmlElementsAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsCustomXmlElementsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsCustomXmlElementsAPI {
	a := &EncodingManifestsDashPeriodsCustomXmlElementsAPI{apiClient: apiClient}
	return a
}

// Create Add Custom XML Element to Period
func (api *EncodingManifestsDashPeriodsCustomXmlElementsAPI) Create(manifestId string, periodId string, customXmlElement model.CustomXmlElement) (*model.CustomXmlElement, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
	}

	var responseModel model.CustomXmlElement
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements", &customXmlElement, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Custom XML Element
func (api *EncodingManifestsDashPeriodsCustomXmlElementsAPI) Delete(manifestId string, periodId string, customXmlElementId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["custom_xml_element_id"] = customXmlElementId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements/{custom_xml_element_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Custom XML Element Details
func (api *EncodingManifestsDashPeriodsCustomXmlElementsAPI) Get(manifestId string, periodId string, customXmlElementId string) (*model.CustomXmlElement, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["custom_xml_element_id"] = customXmlElementId
	}

	var responseModel model.CustomXmlElement
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements/{custom_xml_element_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Custom XML Elements of Period
func (api *EncodingManifestsDashPeriodsCustomXmlElementsAPI) List(manifestId string, periodId string, queryParams ...func(*EncodingManifestsDashPeriodsCustomXmlElementsAPIListQueryParams)) (*pagination.CustomXmlElementsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsCustomXmlElementsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.CustomXmlElementsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsCustomXmlElementsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsCustomXmlElementsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsCustomXmlElementsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
