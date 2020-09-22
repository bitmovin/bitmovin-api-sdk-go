package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveTsAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts' endpoints
type EncodingEncodingsMuxingsProgressiveTsAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsProgressiveTsCustomdataAPI
	// Information communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/information' endpoints
	Information *EncodingEncodingsMuxingsProgressiveTsInformationAPI
	// Id3 communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3' endpoints
	Id3 *EncodingEncodingsMuxingsProgressiveTsId3API
	// Drm communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm' endpoints
	Drm *EncodingEncodingsMuxingsProgressiveTsDrmAPI
}

// NewEncodingEncodingsMuxingsProgressiveTsAPI constructor for EncodingEncodingsMuxingsProgressiveTsAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveTsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsAPI {
	a := &EncodingEncodingsMuxingsProgressiveTsAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsProgressiveTsCustomdataAPIWithClient(apiClient)
	a.Information = NewEncodingEncodingsMuxingsProgressiveTsInformationAPIWithClient(apiClient)
	a.Id3 = NewEncodingEncodingsMuxingsProgressiveTsId3APIWithClient(apiClient)
	a.Drm = NewEncodingEncodingsMuxingsProgressiveTsDrmAPIWithClient(apiClient)

	return a
}

// Create Add Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsAPI) Create(encodingId string, progressiveTsMuxing model.ProgressiveTsMuxing) (*model.ProgressiveTsMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.ProgressiveTsMuxing
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts", &progressiveTsMuxing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Progressive TS muxing details
func (api *EncodingEncodingsMuxingsProgressiveTsAPI) Get(encodingId string, muxingId string) (*model.ProgressiveTsMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.ProgressiveTsMuxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Progressive TS muxings
func (api *EncodingEncodingsMuxingsProgressiveTsAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveTsAPIListQueryParams)) (*pagination.ProgressiveTsMuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsProgressiveTsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ProgressiveTsMuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveTsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveTsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveTsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
