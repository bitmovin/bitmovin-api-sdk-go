package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke' endpoints
type EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI constructor for EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI {
	a := &EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add SPEKE DRM key provider to Progressive TS
func (api *EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI) Create(encodingId string, muxingId string, spekeDrm model.SpekeDrm) (*model.SpekeDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.SpekeDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke", &spekeDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete SPEKE DRM from a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get SPEKE DRM Details of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI) Get(encodingId string, muxingId string, drmId string) (*model.SpekeDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.SpekeDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List SPEKE DRM of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPIListQueryParams)) (*pagination.SpekeDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SpekeDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPIListQueryParams struct {
	Offset string `query:"offset"`
	Limit  string `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
