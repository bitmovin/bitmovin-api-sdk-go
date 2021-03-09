package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsPackedAudioAPI communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio' endpoints
type EncodingEncodingsMuxingsPackedAudioAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsPackedAudioCustomdataAPI
	// Information communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/information' endpoints
	Information *EncodingEncodingsMuxingsPackedAudioInformationAPI
	// Drm intermediary API object with no endpoints
	Drm *EncodingEncodingsMuxingsPackedAudioDrmAPI
}

// NewEncodingEncodingsMuxingsPackedAudioAPI constructor for EncodingEncodingsMuxingsPackedAudioAPI that takes options as argument
func NewEncodingEncodingsMuxingsPackedAudioAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsPackedAudioAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsPackedAudioAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsPackedAudioAPIWithClient constructor for EncodingEncodingsMuxingsPackedAudioAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsPackedAudioAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsPackedAudioAPI {
	a := &EncodingEncodingsMuxingsPackedAudioAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsPackedAudioCustomdataAPIWithClient(apiClient)
	a.Information = NewEncodingEncodingsMuxingsPackedAudioInformationAPIWithClient(apiClient)
	a.Drm = NewEncodingEncodingsMuxingsPackedAudioDrmAPIWithClient(apiClient)

	return a
}

// Create Add Packed Audio muxing
func (api *EncodingEncodingsMuxingsPackedAudioAPI) Create(encodingId string, packedAudioMuxing model.PackedAudioMuxing) (*model.PackedAudioMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.PackedAudioMuxing
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/packed-audio", &packedAudioMuxing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Packed Audio muxing
func (api *EncodingEncodingsMuxingsPackedAudioAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Packed Audio muxing details
func (api *EncodingEncodingsMuxingsPackedAudioAPI) Get(encodingId string, muxingId string) (*model.PackedAudioMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.PackedAudioMuxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Packed Audio muxings
func (api *EncodingEncodingsMuxingsPackedAudioAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsPackedAudioAPIListQueryParams)) (*pagination.PackedAudioMuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsPackedAudioAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.PackedAudioMuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/packed-audio", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsPackedAudioAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsPackedAudioAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsPackedAudioAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
