package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsDirectFileUploadAPI communicates with '/encoding/inputs/direct-file-upload' endpoints
type EncodingInputsDirectFileUploadAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/inputs/direct-file-upload/{input_id}/customData' endpoints
	Customdata *EncodingInputsDirectFileUploadCustomdataAPI
}

// NewEncodingInputsDirectFileUploadAPI constructor for EncodingInputsDirectFileUploadAPI that takes options as argument
func NewEncodingInputsDirectFileUploadAPI(options ...apiclient.APIClientOption) (*EncodingInputsDirectFileUploadAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsDirectFileUploadAPIWithClient(apiClient), nil
}

// NewEncodingInputsDirectFileUploadAPIWithClient constructor for EncodingInputsDirectFileUploadAPI that takes an APIClient as argument
func NewEncodingInputsDirectFileUploadAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsDirectFileUploadAPI {
	a := &EncodingInputsDirectFileUploadAPI{apiClient: apiClient}
	a.Customdata = NewEncodingInputsDirectFileUploadCustomdataAPIWithClient(apiClient)

	return a
}

// Create Direct File Upload Input
func (api *EncodingInputsDirectFileUploadAPI) Create(directFileUploadInput model.DirectFileUploadInput) (*model.DirectFileUploadInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.DirectFileUploadInput
	err := api.apiClient.Post("/encoding/inputs/direct-file-upload", &directFileUploadInput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Direct File Upload Input
func (api *EncodingInputsDirectFileUploadAPI) Delete(inputId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/inputs/direct-file-upload/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Direct File Upload Input Details
func (api *EncodingInputsDirectFileUploadAPI) Get(inputId string) (*model.DirectFileUploadInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.DirectFileUploadInput
	err := api.apiClient.Get("/encoding/inputs/direct-file-upload/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Direct File Upload Inputs
func (api *EncodingInputsDirectFileUploadAPI) List(queryParams ...func(*EncodingInputsDirectFileUploadAPIListQueryParams)) (*pagination.DirectFileUploadInputsListPagination, error) {
	queryParameters := &EncodingInputsDirectFileUploadAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DirectFileUploadInputsListPagination
	err := api.apiClient.Get("/encoding/inputs/direct-file-upload", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsDirectFileUploadAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsDirectFileUploadAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsDirectFileUploadAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
