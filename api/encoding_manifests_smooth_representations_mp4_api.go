package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsSmoothRepresentationsMp4API communicates with '/encoding/manifests/smooth/{manifest_id}/representations/mp4' endpoints
type EncodingManifestsSmoothRepresentationsMp4API struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsSmoothRepresentationsMp4API constructor for EncodingManifestsSmoothRepresentationsMp4API that takes options as argument
func NewEncodingManifestsSmoothRepresentationsMp4API(options ...apiclient.APIClientOption) (*EncodingManifestsSmoothRepresentationsMp4API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsSmoothRepresentationsMp4APIWithClient(apiClient), nil
}

// NewEncodingManifestsSmoothRepresentationsMp4APIWithClient constructor for EncodingManifestsSmoothRepresentationsMp4API that takes an APIClient as argument
func NewEncodingManifestsSmoothRepresentationsMp4APIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsSmoothRepresentationsMp4API {
	a := &EncodingManifestsSmoothRepresentationsMp4API{apiClient: apiClient}
	return a
}

// Create Add MP4 Representation to Smooth Streaming Manifest
func (api *EncodingManifestsSmoothRepresentationsMp4API) Create(manifestId string, smoothStreamingRepresentation model.SmoothStreamingRepresentation) (*model.SmoothStreamingRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.SmoothStreamingRepresentation
	err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/representations/mp4", &smoothStreamingRepresentation, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Smooth Streaming MP4 Representation
func (api *EncodingManifestsSmoothRepresentationsMp4API) Delete(manifestId string, representationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/smooth/{manifest_id}/representations/mp4/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Smooth Streaming MP4 Representation Details
func (api *EncodingManifestsSmoothRepresentationsMp4API) Get(manifestId string, representationId string) (*model.SmoothStreamingRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.SmoothStreamingRepresentation
	err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/representations/mp4/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List MP4 Representation
func (api *EncodingManifestsSmoothRepresentationsMp4API) List(manifestId string, queryParams ...func(*EncodingManifestsSmoothRepresentationsMp4APIListQueryParams)) (*pagination.SmoothStreamingRepresentationsListPagination, error) {
	queryParameters := &EncodingManifestsSmoothRepresentationsMp4APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SmoothStreamingRepresentationsListPagination
	err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/representations/mp4", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsSmoothRepresentationsMp4APIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsSmoothRepresentationsMp4APIListQueryParams struct {
	Offset int64 `query:"offset"`
	Limit  int64 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsSmoothRepresentationsMp4APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
