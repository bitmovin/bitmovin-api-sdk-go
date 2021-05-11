package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsImageAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/image' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsImageAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsImageAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsImageAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsImageAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsImageAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsImageAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsImageAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsImageAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsImageAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsImageAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsImageAPI{apiClient: apiClient}
	return a
}

// Create Add Image AdaptationSet
func (api *EncodingManifestsDashPeriodsAdaptationsetsImageAPI) Create(manifestId string, periodId string, imageAdaptationSet model.ImageAdaptationSet) (*model.ImageAdaptationSet, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
	}

	var responseModel model.ImageAdaptationSet
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/image", &imageAdaptationSet, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Image AdaptationSet
func (api *EncodingManifestsDashPeriodsAdaptationsetsImageAPI) Delete(manifestId string, periodId string, adaptationsetId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/image/{adaptationset_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Image AdaptationSet Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsImageAPI) Get(manifestId string, periodId string, adaptationsetId string) (*model.ImageAdaptationSet, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.ImageAdaptationSet
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/image/{adaptationset_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Image AdaptationSets
func (api *EncodingManifestsDashPeriodsAdaptationsetsImageAPI) List(manifestId string, periodId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsImageAPIListQueryParams)) (*pagination.ImageAdaptationSetsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsImageAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ImageAdaptationSetsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/image", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsImageAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsImageAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsImageAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
