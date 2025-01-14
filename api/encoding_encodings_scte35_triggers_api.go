package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsScte35TriggersAPI communicates with '/encoding/encodings/{encoding_id}/scte-35-triggers' endpoints
type EncodingEncodingsScte35TriggersAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsScte35TriggersAPI constructor for EncodingEncodingsScte35TriggersAPI that takes options as argument
func NewEncodingEncodingsScte35TriggersAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsScte35TriggersAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsScte35TriggersAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsScte35TriggersAPIWithClient constructor for EncodingEncodingsScte35TriggersAPI that takes an APIClient as argument
func NewEncodingEncodingsScte35TriggersAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsScte35TriggersAPI {
	a := &EncodingEncodingsScte35TriggersAPI{apiClient: apiClient}
	return a
}

// Create SCTE 35 trigger
func (api *EncodingEncodingsScte35TriggersAPI) Create(encodingId string, scte35Trigger model.Scte35Trigger) (*model.Scte35Trigger, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Scte35Trigger
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/scte-35-triggers", &scte35Trigger, &responseModel, reqParams)
	return &responseModel, err
}

// Delete SCTE 35 trigger
func (api *EncodingEncodingsScte35TriggersAPI) Delete(encodingId string, scte35triggerId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["scte35trigger_id"] = scte35triggerId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/scte-35-triggers/{scte35trigger_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get SCTE 35 trigger Details
func (api *EncodingEncodingsScte35TriggersAPI) Get(encodingId string, scte35triggerId string) (*model.Scte35Trigger, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["scte35trigger_id"] = scte35triggerId
	}

	var responseModel model.Scte35Trigger
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/scte-35-triggers/{scte35trigger_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all SCTE 35 triggers for an encoding
func (api *EncodingEncodingsScte35TriggersAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsScte35TriggersAPIListQueryParams)) (*pagination.Scte35TriggersListPagination, error) {
	queryParameters := &EncodingEncodingsScte35TriggersAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.Scte35TriggersListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/scte-35-triggers", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsScte35TriggersAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsScte35TriggersAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsScte35TriggersAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
