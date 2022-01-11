package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsAPI struct {
	apiClient *apiclient.APIClient

	// Type communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/type' endpoints
	Type *EncodingManifestsDashPeriodsAdaptationsetsTypeAPI
	// Audio communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/audio' endpoints
	Audio *EncodingManifestsDashPeriodsAdaptationsetsAudioAPI
	// Video communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video' endpoints
	Video *EncodingManifestsDashPeriodsAdaptationsetsVideoAPI
	// Subtitle communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle' endpoints
	Subtitle *EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI
	// Image communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/image' endpoints
	Image *EncodingManifestsDashPeriodsAdaptationsetsImageAPI
	// Representations communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations' endpoints
	Representations *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI
	// Contentprotection communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/contentprotection' endpoints
	Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI
}

// NewEncodingManifestsDashPeriodsAdaptationsetsAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsAPI{apiClient: apiClient}
	a.Type = NewEncodingManifestsDashPeriodsAdaptationsetsTypeAPIWithClient(apiClient)
	a.Audio = NewEncodingManifestsDashPeriodsAdaptationsetsAudioAPIWithClient(apiClient)
	a.Video = NewEncodingManifestsDashPeriodsAdaptationsetsVideoAPIWithClient(apiClient)
	a.Subtitle = NewEncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIWithClient(apiClient)
	a.Image = NewEncodingManifestsDashPeriodsAdaptationsetsImageAPIWithClient(apiClient)
	a.Representations = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIWithClient(apiClient)
	a.Contentprotection = NewEncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIWithClient(apiClient)

	return a
}

// List all AdaptationSets
func (api *EncodingManifestsDashPeriodsAdaptationsetsAPI) List(manifestId string, periodId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsAPIListQueryParams)) (*pagination.AdaptationSetsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AdaptationSetsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
