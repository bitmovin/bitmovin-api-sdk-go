package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsWebmDrmSpekeAPI communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke' endpoints
type EncodingEncodingsMuxingsWebmDrmSpekeAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI
}

// NewEncodingEncodingsMuxingsWebmDrmSpekeAPI constructor for EncodingEncodingsMuxingsWebmDrmSpekeAPI that takes options as argument
func NewEncodingEncodingsMuxingsWebmDrmSpekeAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsWebmDrmSpekeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsWebmDrmSpekeAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsWebmDrmSpekeAPIWithClient constructor for EncodingEncodingsMuxingsWebmDrmSpekeAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsWebmDrmSpekeAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsWebmDrmSpekeAPI {
	a := &EncodingEncodingsMuxingsWebmDrmSpekeAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add SPEKE DRM key provider to a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmSpekeAPI) Create(encodingId string, muxingId string, spekeDrm model.SpekeDrm) (*model.SpekeDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.SpekeDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke", &spekeDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete SPEKE DRM from a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmSpekeAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get SPEKE DRM Details of a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmSpekeAPI) Get(encodingId string, muxingId string, drmId string) (*model.SpekeDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.SpekeDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List SPEKE DRM of a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmSpekeAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsWebmDrmSpekeAPIListQueryParams)) (*pagination.SpekeDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsWebmDrmSpekeAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SpekeDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsWebmDrmSpekeAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsWebmDrmSpekeAPIListQueryParams struct {
	Offset string `query:"offset"`
	Limit  string `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsWebmDrmSpekeAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
