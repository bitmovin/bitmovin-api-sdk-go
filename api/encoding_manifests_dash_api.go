package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashAPI communicates with '/encoding/manifests/dash' endpoints
type EncodingManifestsDashAPI struct {
	apiClient *apiclient.APIClient

	// Default communicates with '/encoding/manifests/dash/default' endpoints
	Default *EncodingManifestsDashDefaultAPI
	// Customdata communicates with '/encoding/manifests/dash/{manifest_id}/customData' endpoints
	Customdata *EncodingManifestsDashCustomdataAPI
	// Periods communicates with '/encoding/manifests/dash/{manifest_id}/periods' endpoints
	Periods *EncodingManifestsDashPeriodsAPI
}

// NewEncodingManifestsDashAPI constructor for EncodingManifestsDashAPI that takes options as argument
func NewEncodingManifestsDashAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashAPIWithClient constructor for EncodingManifestsDashAPI that takes an APIClient as argument
func NewEncodingManifestsDashAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashAPI {
	a := &EncodingManifestsDashAPI{apiClient: apiClient}
	a.Default = NewEncodingManifestsDashDefaultAPIWithClient(apiClient)
	a.Customdata = NewEncodingManifestsDashCustomdataAPIWithClient(apiClient)
	a.Periods = NewEncodingManifestsDashPeriodsAPIWithClient(apiClient)

	return a
}

// Create Custom DASH Manifest
// A Custom DASH Manifest gives you full control over its contents. Add Periods, Adaptation Sets, Representations, Content Protections or Custom XML Elements via the respective endpoints. If you need a simpler solution, create a Default Manifest resource instead. See [documentation](https://developer.bitmovin.com/encoding/docs/default-vs-custom-manifest) page for a comparison
func (api *EncodingManifestsDashAPI) Create(dashManifest model.DashManifest) (*model.DashManifest, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.DashManifest
	err := api.apiClient.Post("/encoding/manifests/dash", &dashManifest, &responseModel, reqParams)
	return &responseModel, err
}

// Delete DASH Manifest
func (api *EncodingManifestsDashAPI) Delete(manifestId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get DASH Manifest Details
func (api *EncodingManifestsDashAPI) Get(manifestId string) (*model.DashManifest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.DashManifest
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetStartRequest Manifest Start Details
func (api *EncodingManifestsDashAPI) GetStartRequest(manifestId string) (*model.StartManifestRequest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.StartManifestRequest
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List DASH Manifests
func (api *EncodingManifestsDashAPI) List(queryParams ...func(*EncodingManifestsDashAPIListQueryParams)) (*pagination.DashManifestsListPagination, error) {
	queryParameters := &EncodingManifestsDashAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DashManifestsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start DASH manifest generation
func (api *EncodingManifestsDashAPI) Start(manifestId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start DASH manifest generation
func (api *EncodingManifestsDashAPI) StartWithRequestBody(manifestId string, startManifestRequest model.StartManifestRequest) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/start", &startManifestRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Status DASH manifest generation status
func (api *EncodingManifestsDashAPI) Status(manifestId string) (*model.ModelTask, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.ModelTask
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/status", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Stop DASH manifest generation
func (api *EncodingManifestsDashAPI) Stop(manifestId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/stop", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashAPIListQueryParams struct {
	Offset     int32  `query:"offset"`
	Limit      int32  `query:"limit"`
	Sort       string `query:"sort"`
	EncodingId string `query:"encodingId"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
