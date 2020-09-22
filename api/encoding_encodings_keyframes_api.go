package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsKeyframesAPI communicates with '/encoding/encodings/{encoding_id}/keyframes' endpoints
type EncodingEncodingsKeyframesAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsKeyframesAPI constructor for EncodingEncodingsKeyframesAPI that takes options as argument
func NewEncodingEncodingsKeyframesAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsKeyframesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsKeyframesAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsKeyframesAPIWithClient constructor for EncodingEncodingsKeyframesAPI that takes an APIClient as argument
func NewEncodingEncodingsKeyframesAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsKeyframesAPI {
	a := &EncodingEncodingsKeyframesAPI{apiClient: apiClient}
	return a
}

// Create Keyframes
func (api *EncodingEncodingsKeyframesAPI) Create(encodingId string, keyframe model.Keyframe) (*model.Keyframe, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Keyframe
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/keyframes", &keyframe, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Keyframe
func (api *EncodingEncodingsKeyframesAPI) Delete(encodingId string, keyframeId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["keyframe_id"] = keyframeId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/keyframes/{keyframe_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Keyframe Details
func (api *EncodingEncodingsKeyframesAPI) Get(encodingId string, keyframeId string) (*model.Keyframe, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["keyframe_id"] = keyframeId
	}

	var responseModel model.Keyframe
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/keyframes/{keyframe_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Keyframes
func (api *EncodingEncodingsKeyframesAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsKeyframesAPIListQueryParams)) (*pagination.KeyframesListPagination, error) {
	queryParameters := &EncodingEncodingsKeyframesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.KeyframesListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/keyframes", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsKeyframesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsKeyframesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsKeyframesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
