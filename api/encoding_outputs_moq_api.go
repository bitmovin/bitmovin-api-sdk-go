package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsMoqAPI communicates with '/encoding/outputs/moq' endpoints
type EncodingOutputsMoqAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/outputs/moq/{output_id}/customData' endpoints
	Customdata *EncodingOutputsMoqCustomdataAPI
}

// NewEncodingOutputsMoqAPI constructor for EncodingOutputsMoqAPI that takes options as argument
func NewEncodingOutputsMoqAPI(options ...apiclient.APIClientOption) (*EncodingOutputsMoqAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsMoqAPIWithClient(apiClient), nil
}

// NewEncodingOutputsMoqAPIWithClient constructor for EncodingOutputsMoqAPI that takes an APIClient as argument
func NewEncodingOutputsMoqAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsMoqAPI {
	a := &EncodingOutputsMoqAPI{apiClient: apiClient}
	a.Customdata = NewEncodingOutputsMoqCustomdataAPIWithClient(apiClient)

	return a
}

// Create MoQ Output
func (api *EncodingOutputsMoqAPI) Create(moqOutput model.MoqOutput) (*model.MoqOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.MoqOutput
	err := api.apiClient.Post("/encoding/outputs/moq", &moqOutput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete MoQ Output
func (api *EncodingOutputsMoqAPI) Delete(outputId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/outputs/moq/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get MoQ Output details
func (api *EncodingOutputsMoqAPI) Get(outputId string) (*model.MoqOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.MoqOutput
	err := api.apiClient.Get("/encoding/outputs/moq/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List MoQ Outputs
func (api *EncodingOutputsMoqAPI) List(queryParams ...func(*EncodingOutputsMoqAPIListQueryParams)) (*pagination.MoqOutputsListPagination, error) {
	queryParameters := &EncodingOutputsMoqAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.MoqOutputsListPagination
	err := api.apiClient.Get("/encoding/outputs/moq", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsMoqAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsMoqAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsMoqAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
