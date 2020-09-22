package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/chunked-text' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI{apiClient: apiClient}
	return a
}

// Create Add Chunked Text Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI) Create(manifestId string, periodId string, adaptationsetId string, dashChunkedTextRepresentation model.DashChunkedTextRepresentation) (*model.DashChunkedTextRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.DashChunkedTextRepresentation
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/chunked-text", &dashChunkedTextRepresentation, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Chunked Text Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/chunked-text/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Chunked Text Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashChunkedTextRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.DashChunkedTextRepresentation
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/chunked-text/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Chunked Text Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIListQueryParams)) (*pagination.DashChunkedTextRepresentationsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DashChunkedTextRepresentationsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/chunked-text", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
