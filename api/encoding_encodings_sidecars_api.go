package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsSidecarsAPI communicates with '/encoding/encodings/{encoding_id}/sidecars' endpoints
type EncodingEncodingsSidecarsAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}/customData' endpoints
	Customdata *EncodingEncodingsSidecarsCustomdataAPI
	// Webvtt communicates with '/encoding/encodings/{encoding_id}/sidecars/webvtt' endpoints
	Webvtt *EncodingEncodingsSidecarsWebvttAPI
}

// NewEncodingEncodingsSidecarsAPI constructor for EncodingEncodingsSidecarsAPI that takes options as argument
func NewEncodingEncodingsSidecarsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsSidecarsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsSidecarsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsSidecarsAPIWithClient constructor for EncodingEncodingsSidecarsAPI that takes an APIClient as argument
func NewEncodingEncodingsSidecarsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsSidecarsAPI {
	a := &EncodingEncodingsSidecarsAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsSidecarsCustomdataAPIWithClient(apiClient)
	a.Webvtt = NewEncodingEncodingsSidecarsWebvttAPIWithClient(apiClient)

	return a
}

// Create Add Sidecar
func (api *EncodingEncodingsSidecarsAPI) Create(encodingId string, sidecarFile model.SidecarFile) (*model.SidecarFile, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.SidecarFile
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/sidecars", &sidecarFile, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Sidecar
func (api *EncodingEncodingsSidecarsAPI) Delete(encodingId string, sidecarId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["sidecar_id"] = sidecarId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Sidecar Details
func (api *EncodingEncodingsSidecarsAPI) Get(encodingId string, sidecarId string) (*model.SidecarFile, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["sidecar_id"] = sidecarId
	}

	var responseModel model.SidecarFile
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Sidecars
func (api *EncodingEncodingsSidecarsAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsSidecarsAPIListQueryParams)) (*pagination.SidecarFilesListPagination, error) {
	queryParameters := &EncodingEncodingsSidecarsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SidecarFilesListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsSidecarsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsSidecarsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsSidecarsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
