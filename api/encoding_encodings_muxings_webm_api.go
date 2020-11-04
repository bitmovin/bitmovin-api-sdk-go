package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsWebmAPI communicates with '/encoding/encodings/{encoding_id}/muxings/webm' endpoints
type EncodingEncodingsMuxingsWebmAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsWebmCustomdataAPI
	// Drm communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm' endpoints
	Drm *EncodingEncodingsMuxingsWebmDrmAPI
}

// NewEncodingEncodingsMuxingsWebmAPI constructor for EncodingEncodingsMuxingsWebmAPI that takes options as argument
func NewEncodingEncodingsMuxingsWebmAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsWebmAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsWebmAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsWebmAPIWithClient constructor for EncodingEncodingsMuxingsWebmAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsWebmAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsWebmAPI {
	a := &EncodingEncodingsMuxingsWebmAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsWebmCustomdataAPIWithClient(apiClient)
	a.Drm = NewEncodingEncodingsMuxingsWebmDrmAPIWithClient(apiClient)

	return a
}

// Create Add WebM muxing
func (api *EncodingEncodingsMuxingsWebmAPI) Create(encodingId string, webmMuxing model.WebmMuxing) (*model.WebmMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.WebmMuxing
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/webm", &webmMuxing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete WebM muxing
func (api *EncodingEncodingsMuxingsWebmAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get WebM muxing details
func (api *EncodingEncodingsMuxingsWebmAPI) Get(encodingId string, muxingId string) (*model.WebmMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.WebmMuxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List WebM muxings
func (api *EncodingEncodingsMuxingsWebmAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsWebmAPIListQueryParams)) (*pagination.WebmMuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsWebmAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WebmMuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsWebmAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsWebmAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsWebmAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
