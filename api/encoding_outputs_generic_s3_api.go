package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsGenericS3API communicates with '/encoding/outputs/generic-s3' endpoints
type EncodingOutputsGenericS3API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/outputs/generic-s3/{output_id}/customData' endpoints
	Customdata *EncodingOutputsGenericS3CustomdataAPI
}

// NewEncodingOutputsGenericS3API constructor for EncodingOutputsGenericS3API that takes options as argument
func NewEncodingOutputsGenericS3API(options ...apiclient.APIClientOption) (*EncodingOutputsGenericS3API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsGenericS3APIWithClient(apiClient), nil
}

// NewEncodingOutputsGenericS3APIWithClient constructor for EncodingOutputsGenericS3API that takes an APIClient as argument
func NewEncodingOutputsGenericS3APIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsGenericS3API {
	a := &EncodingOutputsGenericS3API{apiClient: apiClient}
	a.Customdata = NewEncodingOutputsGenericS3CustomdataAPIWithClient(apiClient)

	return a
}

// Create Generic S3 Output
func (api *EncodingOutputsGenericS3API) Create(genericS3Output model.GenericS3Output) (*model.GenericS3Output, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.GenericS3Output
	err := api.apiClient.Post("/encoding/outputs/generic-s3", &genericS3Output, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Generic S3 Output
func (api *EncodingOutputsGenericS3API) Delete(outputId string) (*model.GenericS3Output, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.GenericS3Output
	err := api.apiClient.Delete("/encoding/outputs/generic-s3/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Generic S3 Output Details
func (api *EncodingOutputsGenericS3API) Get(outputId string) (*model.GenericS3Output, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.GenericS3Output
	err := api.apiClient.Get("/encoding/outputs/generic-s3/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Generic S3 Outputs
func (api *EncodingOutputsGenericS3API) List(queryParams ...func(*EncodingOutputsGenericS3APIListQueryParams)) (*pagination.GenericS3OutputsListPagination, error) {
	queryParameters := &EncodingOutputsGenericS3APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.GenericS3OutputsListPagination
	err := api.apiClient.Get("/encoding/outputs/generic-s3", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsGenericS3APIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsGenericS3APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsGenericS3APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
