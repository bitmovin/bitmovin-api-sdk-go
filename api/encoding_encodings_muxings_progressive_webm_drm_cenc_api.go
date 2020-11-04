package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/cenc' endpoints
type EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/cenc/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmCencAPI constructor for EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmCencAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWebmDrmCencAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmCencAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmCencAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI {
	a := &EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add CENC DRM to a Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI) Create(encodingId string, muxingId string, cencDrm model.CencDrm) (*model.CencDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CencDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/cenc", &cencDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete CENC DRM from a Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get CENC DRM Details of a Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI) Get(encodingId string, muxingId string, drmId string) (*model.CencDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CencDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List CENC DRM configurations of Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveWebmDrmCencAPIListQueryParams)) (*pagination.CencDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsProgressiveWebmDrmCencAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.CencDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/cenc", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveWebmDrmCencAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveWebmDrmCencAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveWebmDrmCencAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
