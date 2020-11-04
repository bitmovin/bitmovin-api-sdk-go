package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsAPI communicates with '/encoding/manifests/hls' endpoints
type EncodingManifestsHlsAPI struct {
	apiClient *apiclient.APIClient

	// Default communicates with '/encoding/manifests/hls/default' endpoints
	Default *EncodingManifestsHlsDefaultAPI
	// Customdata communicates with '/encoding/manifests/hls/{manifest_id}/customData' endpoints
	Customdata *EncodingManifestsHlsCustomdataAPI
	// Streams communicates with '/encoding/manifests/hls/{manifest_id}/streams' endpoints
	Streams *EncodingManifestsHlsStreamsAPI
	// Media intermediary API object with no endpoints
	Media *EncodingManifestsHlsMediaAPI
}

// NewEncodingManifestsHlsAPI constructor for EncodingManifestsHlsAPI that takes options as argument
func NewEncodingManifestsHlsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsHlsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsAPIWithClient constructor for EncodingManifestsHlsAPI that takes an APIClient as argument
func NewEncodingManifestsHlsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsAPI {
	a := &EncodingManifestsHlsAPI{apiClient: apiClient}
	a.Default = NewEncodingManifestsHlsDefaultAPIWithClient(apiClient)
	a.Customdata = NewEncodingManifestsHlsCustomdataAPIWithClient(apiClient)
	a.Streams = NewEncodingManifestsHlsStreamsAPIWithClient(apiClient)
	a.Media = NewEncodingManifestsHlsMediaAPIWithClient(apiClient)

	return a
}

// Create HLS Manifest
func (api *EncodingManifestsHlsAPI) Create(hlsManifest model.HlsManifest) (*model.HlsManifest, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.HlsManifest
	err := api.apiClient.Post("/encoding/manifests/hls", &hlsManifest, &responseModel, reqParams)
	return &responseModel, err
}

// Delete HLS Manifest
func (api *EncodingManifestsHlsAPI) Delete(manifestId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get HLS Manifest Details
func (api *EncodingManifestsHlsAPI) Get(manifestId string) (*model.HlsManifest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.HlsManifest
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List HLS Manifests
func (api *EncodingManifestsHlsAPI) List(queryParams ...func(*EncodingManifestsHlsAPIListQueryParams)) (*pagination.HlsManifestsListPagination, error) {
	queryParameters := &EncodingManifestsHlsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.HlsManifestsListPagination
	err := api.apiClient.Get("/encoding/manifests/hls", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start HLS Manifest Creation
func (api *EncodingManifestsHlsAPI) Start(manifestId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Status HLS Manifest Creation Status
func (api *EncodingManifestsHlsAPI) Status(manifestId string) (*model.ModelTask, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.ModelTask
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/status", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Stop HLS Manifest Creation
func (api *EncodingManifestsHlsAPI) Stop(manifestId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/stop", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsHlsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsAPIListQueryParams struct {
	Offset     int32  `query:"offset"`
	Limit      int32  `query:"limit"`
	EncodingId string `query:"encodingId"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
