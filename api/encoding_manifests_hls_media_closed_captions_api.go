package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsMediaClosedCaptionsAPI communicates with '/encoding/manifests/hls/{manifest_id}/media/closed-captions' endpoints
type EncodingManifestsHlsMediaClosedCaptionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsMediaClosedCaptionsAPI constructor for EncodingManifestsHlsMediaClosedCaptionsAPI that takes options as argument
func NewEncodingManifestsHlsMediaClosedCaptionsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsMediaClosedCaptionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsHlsMediaClosedCaptionsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsMediaClosedCaptionsAPIWithClient constructor for EncodingManifestsHlsMediaClosedCaptionsAPI that takes an APIClient as argument
func NewEncodingManifestsHlsMediaClosedCaptionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsMediaClosedCaptionsAPI {
	a := &EncodingManifestsHlsMediaClosedCaptionsAPI{apiClient: apiClient}
	return a
}

// Create Add Closed Captions Media
func (api *EncodingManifestsHlsMediaClosedCaptionsAPI) Create(manifestId string, closedCaptionsMediaInfo model.ClosedCaptionsMediaInfo) (*model.ClosedCaptionsMediaInfo, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.ClosedCaptionsMediaInfo
	err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/closed-captions", &closedCaptionsMediaInfo, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Closed Captions Media
func (api *EncodingManifestsHlsMediaClosedCaptionsAPI) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["media_id"] = mediaId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/closed-captions/{media_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Closed Captions Media Details
func (api *EncodingManifestsHlsMediaClosedCaptionsAPI) Get(manifestId string, mediaId string) (*model.ClosedCaptionsMediaInfo, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["media_id"] = mediaId
	}

	var responseModel model.ClosedCaptionsMediaInfo
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/closed-captions/{media_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Closed Captions Media
func (api *EncodingManifestsHlsMediaClosedCaptionsAPI) List(manifestId string, queryParams ...func(*EncodingManifestsHlsMediaClosedCaptionsAPIListQueryParams)) (*pagination.ClosedCaptionsMediaInfosListPagination, error) {
	queryParameters := &EncodingManifestsHlsMediaClosedCaptionsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ClosedCaptionsMediaInfosListPagination
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/closed-captions", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsHlsMediaClosedCaptionsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsMediaClosedCaptionsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsMediaClosedCaptionsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
