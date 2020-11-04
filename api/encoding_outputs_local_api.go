package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsLocalAPI communicates with '/encoding/outputs/local' endpoints
type EncodingOutputsLocalAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/outputs/local/{output_id}/customData' endpoints
	Customdata *EncodingOutputsLocalCustomdataAPI
}

// NewEncodingOutputsLocalAPI constructor for EncodingOutputsLocalAPI that takes options as argument
func NewEncodingOutputsLocalAPI(options ...apiclient.APIClientOption) (*EncodingOutputsLocalAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsLocalAPIWithClient(apiClient), nil
}

// NewEncodingOutputsLocalAPIWithClient constructor for EncodingOutputsLocalAPI that takes an APIClient as argument
func NewEncodingOutputsLocalAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsLocalAPI {
	a := &EncodingOutputsLocalAPI{apiClient: apiClient}
	a.Customdata = NewEncodingOutputsLocalCustomdataAPIWithClient(apiClient)

	return a
}

// Create Local Output
func (api *EncodingOutputsLocalAPI) Create(localOutput model.LocalOutput) (*model.LocalOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.LocalOutput
	err := api.apiClient.Post("/encoding/outputs/local", &localOutput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Local Output
func (api *EncodingOutputsLocalAPI) Delete(outputId string) (*model.LocalOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.LocalOutput
	err := api.apiClient.Delete("/encoding/outputs/local/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Local Output Details
func (api *EncodingOutputsLocalAPI) Get(outputId string) (*model.LocalOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.LocalOutput
	err := api.apiClient.Get("/encoding/outputs/local/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Local Outputs
func (api *EncodingOutputsLocalAPI) List(queryParams ...func(*EncodingOutputsLocalAPIListQueryParams)) (*pagination.LocalOutputsListPagination, error) {
	queryParameters := &EncodingOutputsLocalAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.LocalOutputsListPagination
	err := api.apiClient.Get("/encoding/outputs/local", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsLocalAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsLocalAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsLocalAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
