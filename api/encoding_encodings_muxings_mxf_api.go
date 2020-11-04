package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMxfAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mxf' endpoints
type EncodingEncodingsMuxingsMxfAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/mxf/{muxing_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsMxfCustomdataAPI
}

// NewEncodingEncodingsMuxingsMxfAPI constructor for EncodingEncodingsMuxingsMxfAPI that takes options as argument
func NewEncodingEncodingsMuxingsMxfAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMxfAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMxfAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMxfAPIWithClient constructor for EncodingEncodingsMuxingsMxfAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMxfAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMxfAPI {
	a := &EncodingEncodingsMuxingsMxfAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsMxfCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add MXF muxing
func (api *EncodingEncodingsMuxingsMxfAPI) Create(encodingId string, mxfMuxing model.MxfMuxing) (*model.MxfMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.MxfMuxing
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mxf", &mxfMuxing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete MXF muxing
func (api *EncodingEncodingsMuxingsMxfAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mxf/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get MXF muxing details
func (api *EncodingEncodingsMuxingsMxfAPI) Get(encodingId string, muxingId string) (*model.MxfMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.MxfMuxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mxf/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List MXF muxings
func (api *EncodingEncodingsMuxingsMxfAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsMxfAPIListQueryParams)) (*pagination.MxfMuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsMxfAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.MxfMuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mxf", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsMxfAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsMxfAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsMxfAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
