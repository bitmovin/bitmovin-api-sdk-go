package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveWavAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-wav' endpoints
type EncodingEncodingsMuxingsProgressiveWavAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-wav/{muxing_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsProgressiveWavCustomdataAPI
	// Information communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-wav/{muxing_id}/information' endpoints
	Information *EncodingEncodingsMuxingsProgressiveWavInformationAPI
}

// NewEncodingEncodingsMuxingsProgressiveWavAPI constructor for EncodingEncodingsMuxingsProgressiveWavAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWavAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWavAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWavAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWavAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWavAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWavAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWavAPI {
	a := &EncodingEncodingsMuxingsProgressiveWavAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsProgressiveWavCustomdataAPIWithClient(apiClient)
	a.Information = NewEncodingEncodingsMuxingsProgressiveWavInformationAPIWithClient(apiClient)

	return a
}

// Create Add Progressive Wav muxing
func (api *EncodingEncodingsMuxingsProgressiveWavAPI) Create(encodingId string, progressiveWavMuxing model.ProgressiveWavMuxing) (*model.ProgressiveWavMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.ProgressiveWavMuxing
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-wav", &progressiveWavMuxing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Progressive WAV muxing
func (api *EncodingEncodingsMuxingsProgressiveWavAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-wav/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Progressive WAV muxing details
func (api *EncodingEncodingsMuxingsProgressiveWavAPI) Get(encodingId string, muxingId string) (*model.ProgressiveWavMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.ProgressiveWavMuxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-wav/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Progressive WAV muxings
func (api *EncodingEncodingsMuxingsProgressiveWavAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveWavAPIListQueryParams)) (*pagination.ProgressiveWavMuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsProgressiveWavAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ProgressiveWavMuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-wav", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveWavAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveWavAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveWavAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
