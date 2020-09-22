package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMp4DrmClearkeyAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey' endpoints
type EncodingEncodingsMuxingsMp4DrmClearkeyAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI
}

// NewEncodingEncodingsMuxingsMp4DrmClearkeyAPI constructor for EncodingEncodingsMuxingsMp4DrmClearkeyAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmClearkeyAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmClearkeyAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMp4DrmClearkeyAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmClearkeyAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmClearkeyAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmClearkeyAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmClearkeyAPI {
	a := &EncodingEncodingsMuxingsMp4DrmClearkeyAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add ClearKey DRM to an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmClearkeyAPI) Create(encodingId string, muxingId string, clearKeyDrm model.ClearKeyDrm) (*model.ClearKeyDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.ClearKeyDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey", &clearKeyDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete ClearKey DRM from an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmClearkeyAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get ClearKey DRM Details of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmClearkeyAPI) Get(encodingId string, muxingId string, drmId string) (*model.ClearKeyDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.ClearKeyDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List ClearKey DRMs of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmClearkeyAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsMp4DrmClearkeyAPIListQueryParams)) (*pagination.ClearKeyDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsMp4DrmClearkeyAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ClearKeyDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsMp4DrmClearkeyAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsMp4DrmClearkeyAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsMp4DrmClearkeyAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
