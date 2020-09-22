package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI struct {
	apiClient *apiclient.APIClient

	// Contentprotection communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}/contentprotection' endpoints
	Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPI
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI{apiClient: apiClient}
	a.Contentprotection = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionAPIWithClient(apiClient)

	return a
}

// Create Add WebM Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI) Create(manifestId string, periodId string, adaptationsetId string, dashWebmRepresentation model.DashWebmRepresentation) (*model.DashWebmRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.DashWebmRepresentation
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm", &dashWebmRepresentation, &responseModel, reqParams)
	return &responseModel, err
}

// Delete WebM Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get WebM Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashWebmRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.DashWebmRepresentation
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all WebM Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIListQueryParams)) (*pagination.DashWebmRepresentationsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DashWebmRepresentationsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
