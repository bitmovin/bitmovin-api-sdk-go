package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsCmafAPI communicates with '/encoding/encodings/{encoding_id}/muxings/cmaf' endpoints
type EncodingEncodingsMuxingsCmafAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsCmafCustomdataAPI
}

// NewEncodingEncodingsMuxingsCmafAPI constructor for EncodingEncodingsMuxingsCmafAPI that takes options as argument
func NewEncodingEncodingsMuxingsCmafAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsCmafAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsCmafAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsCmafAPIWithClient constructor for EncodingEncodingsMuxingsCmafAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsCmafAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsCmafAPI {
	a := &EncodingEncodingsMuxingsCmafAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsCmafCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add CMAF muxing
func (api *EncodingEncodingsMuxingsCmafAPI) Create(encodingId string, cmafMuxing model.CmafMuxing) (*model.CmafMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.CmafMuxing
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/cmaf", &cmafMuxing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete CMAF muxing
func (api *EncodingEncodingsMuxingsCmafAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get CMAF muxing details
func (api *EncodingEncodingsMuxingsCmafAPI) Get(encodingId string, muxingId string) (*model.CmafMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CmafMuxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List CMAF muxings
func (api *EncodingEncodingsMuxingsCmafAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsCmafAPIListQueryParams)) (*pagination.CmafMuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsCmafAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.CmafMuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsCmafAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsCmafAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsCmafAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
