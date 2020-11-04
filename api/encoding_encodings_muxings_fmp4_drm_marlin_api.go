package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmMarlinAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/marlin' endpoints
type EncodingEncodingsMuxingsFmp4DrmMarlinAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/marlin/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI
}

// NewEncodingEncodingsMuxingsFmp4DrmMarlinAPI constructor for EncodingEncodingsMuxingsFmp4DrmMarlinAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmMarlinAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmMarlinAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4DrmMarlinAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmMarlinAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmMarlinAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmMarlinAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmMarlinAPI {
	a := &EncodingEncodingsMuxingsFmp4DrmMarlinAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add Marlin DRM to an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmMarlinAPI) Create(encodingId string, muxingId string, marlinDrm model.MarlinDrm) (*model.MarlinDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.MarlinDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/marlin", &marlinDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Marlin DRM from an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmMarlinAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/marlin/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Marlin DRM Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmMarlinAPI) Get(encodingId string, muxingId string, drmId string) (*model.MarlinDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.MarlinDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/marlin/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Marlin DRMs of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmMarlinAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4DrmMarlinAPIListQueryParams)) (*pagination.MarlinDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsFmp4DrmMarlinAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.MarlinDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/marlin", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4DrmMarlinAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4DrmMarlinAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4DrmMarlinAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
