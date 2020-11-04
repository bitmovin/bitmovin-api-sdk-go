package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay' endpoints
type EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI constructor for EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI {
	a := &EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add FairPlay DRM to a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI) Create(encodingId string, muxingId string, fairPlayDrm model.FairPlayDrm) (*model.FairPlayDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.FairPlayDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay", &fairPlayDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete FairPlay DRM from a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get FairPlay DRM Details of a Progressive TS
func (api *EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI) Get(encodingId string, muxingId string, drmId string) (*model.FairPlayDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.FairPlayDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List FairPlay DRMs of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPIListQueryParams)) (*pagination.FairPlayDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.FairPlayDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
