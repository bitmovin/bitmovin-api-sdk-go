package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMp4DrmMarlinAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin' endpoints
type EncodingEncodingsMuxingsMp4DrmMarlinAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI
}

// NewEncodingEncodingsMuxingsMp4DrmMarlinAPI constructor for EncodingEncodingsMuxingsMp4DrmMarlinAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmMarlinAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmMarlinAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMp4DrmMarlinAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmMarlinAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmMarlinAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmMarlinAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmMarlinAPI {
	a := &EncodingEncodingsMuxingsMp4DrmMarlinAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add Marlin DRM to an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmMarlinAPI) Create(encodingId string, muxingId string, marlinDrm model.MarlinDrm) (*model.MarlinDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.MarlinDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin", &marlinDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Marlin DRM from an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmMarlinAPI) Delete(encodingId string, muxingId string, drmId string) (*model.MarlinDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.MarlinDrm
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Marlin DRM Details of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmMarlinAPI) Get(encodingId string, muxingId string, drmId string) (*model.MarlinDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.MarlinDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Marlin DRMs of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmMarlinAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsMp4DrmMarlinAPIListQueryParams)) (*pagination.MarlinDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsMp4DrmMarlinAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.MarlinDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsMp4DrmMarlinAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsMp4DrmMarlinAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsMp4DrmMarlinAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
