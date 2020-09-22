package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsVideoAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsVideoAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsVideoAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsVideoAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsVideoAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsVideoAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsVideoAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsVideoAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsVideoAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsVideoAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsVideoAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsVideoAPI{apiClient: apiClient}
	return a
}

// Create Add Video AdaptationSet
func (api *EncodingManifestsDashPeriodsAdaptationsetsVideoAPI) Create(manifestId string, periodId string, videoAdaptationSet model.VideoAdaptationSet) (*model.VideoAdaptationSet, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
	}

	var responseModel model.VideoAdaptationSet
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video", &videoAdaptationSet, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Video AdaptationSet
func (api *EncodingManifestsDashPeriodsAdaptationsetsVideoAPI) Delete(manifestId string, periodId string, adaptationsetId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video/{adaptationset_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Video AdaptationSet Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsVideoAPI) Get(manifestId string, periodId string, adaptationsetId string) (*model.VideoAdaptationSet, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.VideoAdaptationSet
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video/{adaptationset_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Video AdaptationSets
func (api *EncodingManifestsDashPeriodsAdaptationsetsVideoAPI) List(manifestId string, periodId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsVideoAPIListQueryParams)) (*pagination.VideoAdaptationSetsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsVideoAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.VideoAdaptationSetsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsVideoAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsVideoAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsVideoAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
