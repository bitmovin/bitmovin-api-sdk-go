package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsSmoothAPI communicates with '/encoding/manifests/smooth' endpoints
type EncodingManifestsSmoothAPI struct {
	apiClient *apiclient.APIClient

	// Default communicates with '/encoding/manifests/smooth/default' endpoints
	Default *EncodingManifestsSmoothDefaultAPI
	// Customdata communicates with '/encoding/manifests/smooth/{manifest_id}/customData' endpoints
	Customdata *EncodingManifestsSmoothCustomdataAPI
	// Representations intermediary API object with no endpoints
	Representations *EncodingManifestsSmoothRepresentationsAPI
	// Contentprotection communicates with '/encoding/manifests/smooth/{manifest_id}/contentprotection' endpoints
	Contentprotection *EncodingManifestsSmoothContentprotectionAPI
}

// NewEncodingManifestsSmoothAPI constructor for EncodingManifestsSmoothAPI that takes options as argument
func NewEncodingManifestsSmoothAPI(options ...apiclient.APIClientOption) (*EncodingManifestsSmoothAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsSmoothAPIWithClient(apiClient), nil
}

// NewEncodingManifestsSmoothAPIWithClient constructor for EncodingManifestsSmoothAPI that takes an APIClient as argument
func NewEncodingManifestsSmoothAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsSmoothAPI {
	a := &EncodingManifestsSmoothAPI{apiClient: apiClient}
	a.Default = NewEncodingManifestsSmoothDefaultAPIWithClient(apiClient)
	a.Customdata = NewEncodingManifestsSmoothCustomdataAPIWithClient(apiClient)
	a.Representations = NewEncodingManifestsSmoothRepresentationsAPIWithClient(apiClient)
	a.Contentprotection = NewEncodingManifestsSmoothContentprotectionAPIWithClient(apiClient)

	return a
}

// Create Smooth Streaming Manifest
func (api *EncodingManifestsSmoothAPI) Create(smoothStreamingManifest model.SmoothStreamingManifest) (*model.SmoothStreamingManifest, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.SmoothStreamingManifest
	err := api.apiClient.Post("/encoding/manifests/smooth", &smoothStreamingManifest, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Smooth Streaming Manifest
func (api *EncodingManifestsSmoothAPI) Delete(manifestId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/smooth/{manifest_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Smooth Streaming Manifest Details
func (api *EncodingManifestsSmoothAPI) Get(manifestId string) (*model.SmoothStreamingManifest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.SmoothStreamingManifest
	err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Smooth Streaming Manifests
func (api *EncodingManifestsSmoothAPI) List(queryParams ...func(*EncodingManifestsSmoothAPIListQueryParams)) (*pagination.SmoothStreamingManifestsListPagination, error) {
	queryParameters := &EncodingManifestsSmoothAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SmoothStreamingManifestsListPagination
	err := api.apiClient.Get("/encoding/manifests/smooth", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start Smooth Streaming Manifest Creation
func (api *EncodingManifestsSmoothAPI) Start(manifestId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start Smooth Streaming Manifest Creation
func (api *EncodingManifestsSmoothAPI) StartWithRequestBody(manifestId string, body map[string]interface{}) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/start", &body, &responseModel, reqParams)
	return &responseModel, err
}

// Status Smooth Streaming Manifest Creation Status
func (api *EncodingManifestsSmoothAPI) Status(manifestId string) (*model.ModelTask, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.ModelTask
	err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/status", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Stop Smooth Streaming Manifest Creation
func (api *EncodingManifestsSmoothAPI) Stop(manifestId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/stop", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsSmoothAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsSmoothAPIListQueryParams struct {
	Offset     int32  `query:"offset"`
	Limit      int32  `query:"limit"`
	EncodingId string `query:"encodingId"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsSmoothAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
