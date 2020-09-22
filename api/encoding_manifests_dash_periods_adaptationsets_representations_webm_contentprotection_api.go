package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}/contentprotection' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI{apiClient: apiClient}
	return a
}

// Create Add Content Protection to WebM Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI) Create(manifestId string, periodId string, adaptationsetId string, representationId string, contentProtection model.ContentProtection) (*model.ContentProtection, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.ContentProtection
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}/contentprotection", &contentProtection, &responseModel, reqParams)
	return &responseModel, err
}

// Delete WebM Representation Content Protection
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.PathParams["contentprotection_id"] = contentprotectionId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get WebM Representation Content Protection Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.ContentProtection, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.PathParams["contentprotection_id"] = contentprotectionId
	}

	var responseModel model.ContentProtection
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all WebM Representation Content Protections
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI) List(manifestId string, periodId string, adaptationsetId string, representationId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPIListQueryParams)) (*pagination.ContentProtectionsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPIListQueryParams{}
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
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}/contentprotection", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
