package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}/contentprotection' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI{apiClient: apiClient}
	return a
}

// Create Add Content Protection to DRM fMP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI) Create(manifestId string, periodId string, adaptationsetId string, representationId string, contentProtection model.ContentProtection) (*model.ContentProtection, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.ContentProtection
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}/contentprotection", &contentProtection, &responseModel, reqParams)
	return &responseModel, err
}

// Delete DRM fMP4 Representation Content Protection
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.PathParams["contentprotection_id"] = contentprotectionId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get DRM fMP4 Representation Content Protection Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.ContentProtection, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.PathParams["contentprotection_id"] = contentprotectionId
	}

	var responseModel model.ContentProtection
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all DRM fMP4 Representation Content Protections
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI) List(manifestId string, periodId string, adaptationsetId string, representationId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPIListQueryParams)) (*pagination.ContentProtectionsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPIListQueryParams{}
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
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}/contentprotection", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
