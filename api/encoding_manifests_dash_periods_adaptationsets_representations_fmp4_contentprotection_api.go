package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI{apiClient: apiClient}
	return a
}

// Create Add Content Protection to fMP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI) Create(manifestId string, periodId string, adaptationsetId string, representationId string, contentProtection model.ContentProtection) (*model.ContentProtection, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.ContentProtection
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection", &contentProtection, &responseModel, reqParams)
	return &responseModel, err
}

// Delete fMP4 Representation Content Protection
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.PathParams["contentprotection_id"] = contentprotectionId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get fMP4 Representation Content Protection Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.ContentProtection, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.PathParams["contentprotection_id"] = contentprotectionId
	}

	var responseModel model.ContentProtection
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all fMP4 Representation Content Protections
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI) List(manifestId string, periodId string, adaptationsetId string, representationId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPIListQueryParams)) (*pagination.ContentProtectionsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ContentProtectionsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
