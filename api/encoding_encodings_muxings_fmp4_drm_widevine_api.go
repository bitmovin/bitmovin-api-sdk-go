package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmWidevineAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/widevine' endpoints
type EncodingEncodingsMuxingsFmp4DrmWidevineAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/widevine/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI
}

// NewEncodingEncodingsMuxingsFmp4DrmWidevineAPI constructor for EncodingEncodingsMuxingsFmp4DrmWidevineAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmWidevineAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmWidevineAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4DrmWidevineAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmWidevineAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmWidevineAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmWidevineAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmWidevineAPI {
	a := &EncodingEncodingsMuxingsFmp4DrmWidevineAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add Widevine DRM to an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmWidevineAPI) Create(encodingId string, muxingId string, widevineDrm model.WidevineDrm) (*model.WidevineDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.WidevineDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/widevine", &widevineDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Widevine DRM from an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmWidevineAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/widevine/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Widevine DRM Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmWidevineAPI) Get(encodingId string, muxingId string, drmId string) (*model.WidevineDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.WidevineDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/widevine/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Widevine DRMs of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmWidevineAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4DrmWidevineAPIListQueryParams)) (*pagination.WidevineDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsFmp4DrmWidevineAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WidevineDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/widevine", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4DrmWidevineAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4DrmWidevineAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4DrmWidevineAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
