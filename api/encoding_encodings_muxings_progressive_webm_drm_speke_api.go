package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/speke' endpoints
type EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI constructor for EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI {
	a := &EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add SPEKE DRM key provider to Progressive WebM
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI) Create(encodingId string, muxingId string, spekeDrm model.SpekeDrm) (*model.SpekeDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.SpekeDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/speke", &spekeDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete SPEKE DRM from a Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get SPEKE DRM Details of a Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI) Get(encodingId string, muxingId string, drmId string) (*model.SpekeDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.SpekeDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List SPEKE DRM of a Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPIListQueryParams)) (*pagination.SpekeDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SpekeDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/speke", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPIListQueryParams struct {
	Offset string `query:"offset"`
	Limit  string `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
