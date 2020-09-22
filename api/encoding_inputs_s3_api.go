package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsS3API communicates with '/encoding/inputs/s3' endpoints
type EncodingInputsS3API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/inputs/s3/{input_id}/customData' endpoints
	Customdata *EncodingInputsS3CustomdataAPI
}

// NewEncodingInputsS3API constructor for EncodingInputsS3API that takes options as argument
func NewEncodingInputsS3API(options ...apiclient.APIClientOption) (*EncodingInputsS3API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsS3APIWithClient(apiClient), nil
}

// NewEncodingInputsS3APIWithClient constructor for EncodingInputsS3API that takes an APIClient as argument
func NewEncodingInputsS3APIWithClient(apiClient *apiclient.APIClient) *EncodingInputsS3API {
	a := &EncodingInputsS3API{apiClient: apiClient}
	a.Customdata = NewEncodingInputsS3CustomdataAPIWithClient(apiClient)

	return a
}

// Create S3 Input
func (api *EncodingInputsS3API) Create(s3Input model.S3Input) (*model.S3Input, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.S3Input
	err := api.apiClient.Post("/encoding/inputs/s3", &s3Input, &responseModel, reqParams)
	return &responseModel, err
}

// Delete S3 Input
func (api *EncodingInputsS3API) Delete(inputId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/inputs/s3/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get S3 Input Details
func (api *EncodingInputsS3API) Get(inputId string) (*model.S3Input, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.S3Input
	err := api.apiClient.Get("/encoding/inputs/s3/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List S3 Inputs
func (api *EncodingInputsS3API) List(queryParams ...func(*EncodingInputsS3APIListQueryParams)) (*pagination.S3InputsListPagination, error) {
	queryParameters := &EncodingInputsS3APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.S3InputsListPagination
	err := api.apiClient.Get("/encoding/inputs/s3", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsS3APIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsS3APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsS3APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
