package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsAPI communicates with '/encoding/manifests' endpoints
type EncodingManifestsAPI struct {
	apiClient *apiclient.APIClient

	// Type communicates with '/encoding/manifests/{manifest_id}/type' endpoints
	Type *EncodingManifestsTypeAPI
	// Dash communicates with '/encoding/manifests/dash' endpoints
	Dash *EncodingManifestsDashAPI
	// Hls communicates with '/encoding/manifests/hls' endpoints
	Hls *EncodingManifestsHlsAPI
	// Smooth communicates with '/encoding/manifests/smooth' endpoints
	Smooth *EncodingManifestsSmoothAPI
}

// NewEncodingManifestsAPI constructor for EncodingManifestsAPI that takes options as argument
func NewEncodingManifestsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsAPIWithClient constructor for EncodingManifestsAPI that takes an APIClient as argument
func NewEncodingManifestsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsAPI {
	a := &EncodingManifestsAPI{apiClient: apiClient}
	a.Type = NewEncodingManifestsTypeAPIWithClient(apiClient)
	a.Dash = NewEncodingManifestsDashAPIWithClient(apiClient)
	a.Hls = NewEncodingManifestsHlsAPIWithClient(apiClient)
	a.Smooth = NewEncodingManifestsSmoothAPIWithClient(apiClient)

	return a
}

// List all Manifests
func (api *EncodingManifestsAPI) List(queryParams ...func(*EncodingManifestsAPIListQueryParams)) (*pagination.ManifestsListPagination, error) {
	queryParameters := &EncodingManifestsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ManifestsListPagination
	err := api.apiClient.Get("/encoding/manifests", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Sort   string `query:"sort"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
