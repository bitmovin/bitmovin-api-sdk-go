package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsAkamaiNetstorageAPI communicates with '/encoding/inputs/akamai-netstorage' endpoints
type EncodingInputsAkamaiNetstorageAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/inputs/akamai-netstorage/{input_id}/customData' endpoints
	Customdata *EncodingInputsAkamaiNetstorageCustomdataAPI
}

// NewEncodingInputsAkamaiNetstorageAPI constructor for EncodingInputsAkamaiNetstorageAPI that takes options as argument
func NewEncodingInputsAkamaiNetstorageAPI(options ...apiclient.APIClientOption) (*EncodingInputsAkamaiNetstorageAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsAkamaiNetstorageAPIWithClient(apiClient), nil
}

// NewEncodingInputsAkamaiNetstorageAPIWithClient constructor for EncodingInputsAkamaiNetstorageAPI that takes an APIClient as argument
func NewEncodingInputsAkamaiNetstorageAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsAkamaiNetstorageAPI {
	a := &EncodingInputsAkamaiNetstorageAPI{apiClient: apiClient}
	a.Customdata = NewEncodingInputsAkamaiNetstorageCustomdataAPIWithClient(apiClient)

	return a
}

// Create Akamai NetStorage Input
func (api *EncodingInputsAkamaiNetstorageAPI) Create(akamaiNetStorageInput model.AkamaiNetStorageInput) (*model.AkamaiNetStorageInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AkamaiNetStorageInput
	err := api.apiClient.Post("/encoding/inputs/akamai-netstorage", &akamaiNetStorageInput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Akamai NetStorage Input
func (api *EncodingInputsAkamaiNetstorageAPI) Delete(inputId string) (*model.AkamaiNetStorageInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.AkamaiNetStorageInput
	err := api.apiClient.Delete("/encoding/inputs/akamai-netstorage/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Akamai NetStorage Input Details
func (api *EncodingInputsAkamaiNetstorageAPI) Get(inputId string) (*model.AkamaiNetStorageInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.AkamaiNetStorageInput
	err := api.apiClient.Get("/encoding/inputs/akamai-netstorage/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Akamai NetStorage Inputs
func (api *EncodingInputsAkamaiNetstorageAPI) List(queryParams ...func(*EncodingInputsAkamaiNetstorageAPIListQueryParams)) (*pagination.AkamaiNetStorageInputsListPagination, error) {
	queryParameters := &EncodingInputsAkamaiNetstorageAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AkamaiNetStorageInputsListPagination
	err := api.apiClient.Get("/encoding/inputs/akamai-netstorage", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsAkamaiNetstorageAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsAkamaiNetstorageAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsAkamaiNetstorageAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
