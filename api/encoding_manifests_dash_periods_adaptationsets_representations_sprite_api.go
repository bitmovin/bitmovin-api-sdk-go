package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/sprite' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI{apiClient: apiClient}
	return a
}

// Create Add Sprite Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI) Create(manifestId string, periodId string, adaptationsetId string, spriteRepresentation model.SpriteRepresentation) (*model.SpriteRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.SpriteRepresentation
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/sprite", &spriteRepresentation, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Sprite Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/sprite/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Sprite Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.SpriteRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.SpriteRepresentation
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/sprite/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Sprite Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIListQueryParams)) (*pagination.SpriteRepresentationsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SpriteRepresentationsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/sprite", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
