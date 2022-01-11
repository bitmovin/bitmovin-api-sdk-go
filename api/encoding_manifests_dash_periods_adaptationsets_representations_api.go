package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI struct {
	apiClient *apiclient.APIClient

	// Type communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/{representation_id}/type' endpoints
	Type *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI
	// Vtt communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/vtt' endpoints
	Vtt *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI
	// Imsc communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/imsc' endpoints
	Imsc *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPI
	// Sprite communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/sprite' endpoints
	Sprite *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI
	// Fmp4 communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4' endpoints
	Fmp4 *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API
	// ChunkedText communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/chunked-text' endpoints
	ChunkedText *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI
	// Cmaf communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf' endpoints
	Cmaf *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI
	// Mp4 communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4' endpoints
	Mp4 *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API
	// Webm communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm' endpoints
	Webm *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI
	// ProgressiveWebm communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/progressive-webm' endpoints
	ProgressiveWebm *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI{apiClient: apiClient}
	a.Type = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPIWithClient(apiClient)
	a.Vtt = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIWithClient(apiClient)
	a.Imsc = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsImscAPIWithClient(apiClient)
	a.Sprite = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIWithClient(apiClient)
	a.Fmp4 = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIWithClient(apiClient)
	a.ChunkedText = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIWithClient(apiClient)
	a.Cmaf = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIWithClient(apiClient)
	a.Mp4 = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIWithClient(apiClient)
	a.Webm = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIWithClient(apiClient)
	a.ProgressiveWebm = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIWithClient(apiClient)

	return a
}

// List all DASH Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIListQueryParams)) (*pagination.DashRepresentationsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DashRepresentationsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
