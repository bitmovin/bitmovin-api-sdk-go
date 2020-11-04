package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsAudioAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/audio' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsAudioAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsAudioAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsAudioAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsAudioAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsAudioAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsAudioAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsAudioAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsAudioAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsAudioAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsAudioAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsAudioAPI{apiClient: apiClient}
	return a
}

// Create Add Audio AdaptationSet
func (api *EncodingManifestsDashPeriodsAdaptationsetsAudioAPI) Create(manifestId string, periodId string, audioAdaptationSet model.AudioAdaptationSet) (*model.AudioAdaptationSet, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
	}

	var responseModel model.AudioAdaptationSet
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/audio", &audioAdaptationSet, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Audio AdaptationSet
func (api *EncodingManifestsDashPeriodsAdaptationsetsAudioAPI) Delete(manifestId string, periodId string, adaptationsetId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/audio/{adaptationset_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Audio AdaptationSet Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsAudioAPI) Get(manifestId string, periodId string, adaptationsetId string) (*model.AudioAdaptationSet, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.AudioAdaptationSet
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/audio/{adaptationset_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Audio AdaptationSets
func (api *EncodingManifestsDashPeriodsAdaptationsetsAudioAPI) List(manifestId string, periodId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsAudioAPIListQueryParams)) (*pagination.AudioAdaptationSetsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsAudioAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AudioAdaptationSetsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/audio", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsAudioAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsAudioAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsAudioAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
