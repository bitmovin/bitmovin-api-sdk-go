package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}/contentprotection' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI{apiClient: apiClient}
	return a
}

// Create Add Content Protection to CMAF Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI) Create(manifestId string, periodId string, adaptationsetId string, representationId string, contentProtection model.ContentProtection) (*model.ContentProtection, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.ContentProtection
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}/contentprotection", &contentProtection, &responseModel, reqParams)
	return &responseModel, err
}

// Delete CMAF Representation Content Protection
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.PathParams["contentprotection_id"] = contentprotectionId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get CMAF Representation Content Protection Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.ContentProtection, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.PathParams["contentprotection_id"] = contentprotectionId
	}

	var responseModel model.ContentProtection
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all CMAF Representation Content Protections
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI) List(manifestId string, periodId string, adaptationsetId string, representationId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPIListQueryParams)) (*pagination.ContentProtectionsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPIListQueryParams{}
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
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}/contentprotection", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
