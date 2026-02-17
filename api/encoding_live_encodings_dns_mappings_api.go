package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingLiveEncodingsDnsMappingsAPI communicates with '/encoding/live/encodings/{encoding_id}/dns-mappings' endpoints
type EncodingLiveEncodingsDnsMappingsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingLiveEncodingsDnsMappingsAPI constructor for EncodingLiveEncodingsDnsMappingsAPI that takes options as argument
func NewEncodingLiveEncodingsDnsMappingsAPI(options ...apiclient.APIClientOption) (*EncodingLiveEncodingsDnsMappingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveEncodingsDnsMappingsAPIWithClient(apiClient), nil
}

// NewEncodingLiveEncodingsDnsMappingsAPIWithClient constructor for EncodingLiveEncodingsDnsMappingsAPI that takes an APIClient as argument
func NewEncodingLiveEncodingsDnsMappingsAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveEncodingsDnsMappingsAPI {
	a := &EncodingLiveEncodingsDnsMappingsAPI{apiClient: apiClient}
	return a
}

// Create new DNS mapping for encoding
func (api *EncodingLiveEncodingsDnsMappingsAPI) Create(encodingId string, dnsMappingRequest model.DnsMappingRequest) (*model.DnsMappingResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.DnsMappingResponse
	err := api.apiClient.Post("/encoding/live/encodings/{encoding_id}/dns-mappings", &dnsMappingRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Delete DNS mapping
func (api *EncodingLiveEncodingsDnsMappingsAPI) Delete(encodingId string, id string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["id"] = id
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/live/encodings/{encoding_id}/dns-mappings/{id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteAll Delete all DNS mappings for encoding
func (api *EncodingLiveEncodingsDnsMappingsAPI) DeleteAll(encodingId string) (*model.BitmovinResponseList, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponseList
	err := api.apiClient.Delete("/encoding/live/encodings/{encoding_id}/dns-mappings", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get DNS mapping details
func (api *EncodingLiveEncodingsDnsMappingsAPI) Get(encodingId string, id string) (*model.DnsMappingResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["id"] = id
	}

	var responseModel model.DnsMappingResponse
	err := api.apiClient.Get("/encoding/live/encodings/{encoding_id}/dns-mappings/{id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List DNS mappings for encoding
func (api *EncodingLiveEncodingsDnsMappingsAPI) List(encodingId string, queryParams ...func(*EncodingLiveEncodingsDnsMappingsAPIListQueryParams)) (*pagination.DnsMappingResponsesListPagination, error) {
	queryParameters := &EncodingLiveEncodingsDnsMappingsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DnsMappingResponsesListPagination
	err := api.apiClient.Get("/encoding/live/encodings/{encoding_id}/dns-mappings", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingLiveEncodingsDnsMappingsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingLiveEncodingsDnsMappingsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingLiveEncodingsDnsMappingsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
