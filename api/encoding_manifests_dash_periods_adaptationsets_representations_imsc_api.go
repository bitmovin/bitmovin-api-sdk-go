package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/imsc' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI{apiClient: apiClient}
	return a
}

// Create Add IMSC Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI) Create(manifestId string, periodId string, adaptationsetId string, dashImscRepresentation model.DashImscRepresentation) (*model.DashImscRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.DashImscRepresentation
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/imsc", &dashImscRepresentation, &responseModel, reqParams)
	return &responseModel, err
}

// Delete IMSC Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/imsc/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get IMSC Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashImscRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.DashImscRepresentation
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/imsc/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all IMSC Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPIListQueryParams)) (*pagination.DashImscRepresentationsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DashImscRepresentationsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/imsc", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
