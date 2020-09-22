package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMp3API communicates with '/encoding/encodings/{encoding_id}/muxings/mp3' endpoints
type EncodingEncodingsMuxingsMp3API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsMp3CustomdataAPI
	// Information communicates with '/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}/information' endpoints
	Information *EncodingEncodingsMuxingsMp3InformationAPI
}

// NewEncodingEncodingsMuxingsMp3API constructor for EncodingEncodingsMuxingsMp3API that takes options as argument
func NewEncodingEncodingsMuxingsMp3API(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp3API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMp3APIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp3APIWithClient constructor for EncodingEncodingsMuxingsMp3API that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp3APIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp3API {
	a := &EncodingEncodingsMuxingsMp3API{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsMp3CustomdataAPIWithClient(apiClient)
	a.Information = NewEncodingEncodingsMuxingsMp3InformationAPIWithClient(apiClient)

	return a
}

// Create Add MP3 muxing
func (api *EncodingEncodingsMuxingsMp3API) Create(encodingId string, mp3Muxing model.Mp3Muxing) (*model.Mp3Muxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Mp3Muxing
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp3", &mp3Muxing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete MP3 muxing
func (api *EncodingEncodingsMuxingsMp3API) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get MP3 muxing details
func (api *EncodingEncodingsMuxingsMp3API) Get(encodingId string, muxingId string) (*model.Mp3Muxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.Mp3Muxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List MP3 muxings
func (api *EncodingEncodingsMuxingsMp3API) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsMp3APIListQueryParams)) (*pagination.Mp3MuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsMp3APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.Mp3MuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp3", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsMp3APIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsMp3APIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsMp3APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
