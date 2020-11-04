package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsWebmDrmCencAPI communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc' endpoints
type EncodingEncodingsMuxingsWebmDrmCencAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsWebmDrmCencCustomdataAPI
}

// NewEncodingEncodingsMuxingsWebmDrmCencAPI constructor for EncodingEncodingsMuxingsWebmDrmCencAPI that takes options as argument
func NewEncodingEncodingsMuxingsWebmDrmCencAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsWebmDrmCencAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsWebmDrmCencAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsWebmDrmCencAPIWithClient constructor for EncodingEncodingsMuxingsWebmDrmCencAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsWebmDrmCencAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsWebmDrmCencAPI {
	a := &EncodingEncodingsMuxingsWebmDrmCencAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsWebmDrmCencCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add CENC DRM to a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmCencAPI) Create(encodingId string, muxingId string, cencDrm model.CencDrm) (*model.CencDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CencDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc", &cencDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete CENC DRM from a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmCencAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get CENC DRM Details of a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmCencAPI) Get(encodingId string, muxingId string, drmId string) (*model.CencDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CencDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List CENC DRMs of a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmCencAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsWebmDrmCencAPIListQueryParams)) (*pagination.CencDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsWebmDrmCencAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.CencDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsWebmDrmCencAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsWebmDrmCencAPIListQueryParams struct {
	Offset string `query:"offset"`
	Limit  string `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsWebmDrmCencAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
