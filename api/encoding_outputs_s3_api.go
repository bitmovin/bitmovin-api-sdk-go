package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsS3API communicates with '/encoding/outputs/s3' endpoints
type EncodingOutputsS3API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/outputs/s3/{output_id}/customData' endpoints
	Customdata *EncodingOutputsS3CustomdataAPI
}

// NewEncodingOutputsS3API constructor for EncodingOutputsS3API that takes options as argument
func NewEncodingOutputsS3API(options ...apiclient.APIClientOption) (*EncodingOutputsS3API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsS3APIWithClient(apiClient), nil
}

// NewEncodingOutputsS3APIWithClient constructor for EncodingOutputsS3API that takes an APIClient as argument
func NewEncodingOutputsS3APIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsS3API {
	a := &EncodingOutputsS3API{apiClient: apiClient}
	a.Customdata = NewEncodingOutputsS3CustomdataAPIWithClient(apiClient)

	return a
}

// Create S3 Output
func (api *EncodingOutputsS3API) Create(s3Output model.S3Output) (*model.S3Output, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.S3Output
	err := api.apiClient.Post("/encoding/outputs/s3", &s3Output, &responseModel, reqParams)
	return &responseModel, err
}

// Delete S3 Output
func (api *EncodingOutputsS3API) Delete(outputId string) (*model.S3Output, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.S3Output
	err := api.apiClient.Delete("/encoding/outputs/s3/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get S3 Output Details
func (api *EncodingOutputsS3API) Get(outputId string) (*model.S3Output, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.S3Output
	err := api.apiClient.Get("/encoding/outputs/s3/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List S3 Outputs
func (api *EncodingOutputsS3API) List(queryParams ...func(*EncodingOutputsS3APIListQueryParams)) (*pagination.S3OutputsListPagination, error) {
	queryParameters := &EncodingOutputsS3APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.S3OutputsListPagination
	err := api.apiClient.Get("/encoding/outputs/s3", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsS3APIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsS3APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsS3APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
