package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveTsId3RawAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw' endpoints
type EncodingEncodingsMuxingsProgressiveTsId3RawAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw/{id3_tag_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI
}

// NewEncodingEncodingsMuxingsProgressiveTsId3RawAPI constructor for EncodingEncodingsMuxingsProgressiveTsId3RawAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3RawAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsId3RawAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveTsId3RawAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsId3RawAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsId3RawAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3RawAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsId3RawAPI {
	a := &EncodingEncodingsMuxingsProgressiveTsId3RawAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add Raw ID3 Tag to a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3RawAPI) Create(encodingId string, muxingId string, rawId3Tag model.RawId3Tag) (*model.RawId3Tag, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.RawId3Tag
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw", &rawId3Tag, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Raw ID3 Tag of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3RawAPI) Delete(encodingId string, muxingId string, id3TagId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["id3_tag_id"] = id3TagId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw/{id3_tag_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Raw ID3 Tag Details of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3RawAPI) Get(encodingId string, muxingId string, id3TagId string) (*model.RawId3Tag, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["id3_tag_id"] = id3TagId
	}

	var responseModel model.RawId3Tag
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw/{id3_tag_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Raw ID3 Tags of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3RawAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveTsId3RawAPIListQueryParams)) (*pagination.RawId3TagsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsProgressiveTsId3RawAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.RawId3TagsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveTsId3RawAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveTsId3RawAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveTsId3RawAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
