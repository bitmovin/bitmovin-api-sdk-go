package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveWebmAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm' endpoints
type EncodingEncodingsMuxingsProgressiveWebmAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsProgressiveWebmCustomdataAPI
	// Information communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/information' endpoints
	Information *EncodingEncodingsMuxingsProgressiveWebmInformationAPI
	// Drm communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm' endpoints
	Drm *EncodingEncodingsMuxingsProgressiveWebmDrmAPI
}

// NewEncodingEncodingsMuxingsProgressiveWebmAPI constructor for EncodingEncodingsMuxingsProgressiveWebmAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWebmAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWebmAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWebmAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWebmAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWebmAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWebmAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWebmAPI {
	a := &EncodingEncodingsMuxingsProgressiveWebmAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsProgressiveWebmCustomdataAPIWithClient(apiClient)
	a.Information = NewEncodingEncodingsMuxingsProgressiveWebmInformationAPIWithClient(apiClient)
	a.Drm = NewEncodingEncodingsMuxingsProgressiveWebmDrmAPIWithClient(apiClient)

	return a
}

// Create Add Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmAPI) Create(encodingId string, progressiveWebmMuxing model.ProgressiveWebmMuxing) (*model.ProgressiveWebmMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.ProgressiveWebmMuxing
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-webm", &progressiveWebmMuxing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Progressive WebM muxing details
func (api *EncodingEncodingsMuxingsProgressiveWebmAPI) Get(encodingId string, muxingId string) (*model.ProgressiveWebmMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.ProgressiveWebmMuxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Progressive WebM muxings
func (api *EncodingEncodingsMuxingsProgressiveWebmAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveWebmAPIListQueryParams)) (*pagination.ProgressiveWebmMuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsProgressiveWebmAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ProgressiveWebmMuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveWebmAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveWebmAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveWebmAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
