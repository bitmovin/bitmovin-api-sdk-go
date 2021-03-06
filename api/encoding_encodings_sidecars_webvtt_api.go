package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsSidecarsWebvttAPI communicates with '/encoding/encodings/{encoding_id}/sidecars/webvtt' endpoints
type EncodingEncodingsSidecarsWebvttAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsSidecarsWebvttAPI constructor for EncodingEncodingsSidecarsWebvttAPI that takes options as argument
func NewEncodingEncodingsSidecarsWebvttAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsSidecarsWebvttAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsSidecarsWebvttAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsSidecarsWebvttAPIWithClient constructor for EncodingEncodingsSidecarsWebvttAPI that takes an APIClient as argument
func NewEncodingEncodingsSidecarsWebvttAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsSidecarsWebvttAPI {
	a := &EncodingEncodingsSidecarsWebvttAPI{apiClient: apiClient}
	return a
}

// Create Add WebVTT sidecar file
func (api *EncodingEncodingsSidecarsWebvttAPI) Create(encodingId string, webVttSidecarFile model.WebVttSidecarFile) (*model.WebVttSidecarFile, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.WebVttSidecarFile
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/sidecars/webvtt", &webVttSidecarFile, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Sidecar
func (api *EncodingEncodingsSidecarsWebvttAPI) Delete(encodingId string, sidecarId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["sidecar_id"] = sidecarId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/sidecars/webvtt/{sidecar_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get WebVTT Sidecar Details
func (api *EncodingEncodingsSidecarsWebvttAPI) Get(encodingId string, sidecarId string) (*model.WebVttSidecarFile, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["sidecar_id"] = sidecarId
	}

	var responseModel model.WebVttSidecarFile
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars/webvtt/{sidecar_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List WebVTT sidecar files
func (api *EncodingEncodingsSidecarsWebvttAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsSidecarsWebvttAPIListQueryParams)) (*pagination.WebVttSidecarFilesListPagination, error) {
	queryParameters := &EncodingEncodingsSidecarsWebvttAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WebVttSidecarFilesListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars/webvtt", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsSidecarsWebvttAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsSidecarsWebvttAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsSidecarsWebvttAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
