package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsTsDrmFairplayAPI communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/fairplay' endpoints
type EncodingEncodingsMuxingsTsDrmFairplayAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/fairplay/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI
}

// NewEncodingEncodingsMuxingsTsDrmFairplayAPI constructor for EncodingEncodingsMuxingsTsDrmFairplayAPI that takes options as argument
func NewEncodingEncodingsMuxingsTsDrmFairplayAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTsDrmFairplayAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsTsDrmFairplayAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTsDrmFairplayAPIWithClient constructor for EncodingEncodingsMuxingsTsDrmFairplayAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTsDrmFairplayAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTsDrmFairplayAPI {
	a := &EncodingEncodingsMuxingsTsDrmFairplayAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsTsDrmFairplayCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add FairPlay DRM to a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmFairplayAPI) Create(encodingId string, muxingId string, fairPlayDrm model.FairPlayDrm) (*model.FairPlayDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.FairPlayDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/fairplay", &fairPlayDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete FairPlay DRM from a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmFairplayAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get FairPlay DRM Details of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmFairplayAPI) Get(encodingId string, muxingId string, drmId string) (*model.FairPlayDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.FairPlayDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List FairPlay DRMs of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmFairplayAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsTsDrmFairplayAPIListQueryParams)) (*pagination.FairPlayDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsTsDrmFairplayAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.FairPlayDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/fairplay", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsTsDrmFairplayAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsTsDrmFairplayAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsTsDrmFairplayAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
