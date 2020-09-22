package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsTcpAPI communicates with '/encoding/inputs/tcp' endpoints
type EncodingInputsTcpAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsTcpAPI constructor for EncodingInputsTcpAPI that takes options as argument
func NewEncodingInputsTcpAPI(options ...apiclient.APIClientOption) (*EncodingInputsTcpAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsTcpAPIWithClient(apiClient), nil
}

// NewEncodingInputsTcpAPIWithClient constructor for EncodingInputsTcpAPI that takes an APIClient as argument
func NewEncodingInputsTcpAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsTcpAPI {
	a := &EncodingInputsTcpAPI{apiClient: apiClient}
	return a
}

// Get TCP Input Details
func (api *EncodingInputsTcpAPI) Get(inputId string) (*model.TcpInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.TcpInput
	err := api.apiClient.Get("/encoding/inputs/tcp/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List TCP inputs
func (api *EncodingInputsTcpAPI) List(queryParams ...func(*EncodingInputsTcpAPIListQueryParams)) (*pagination.TcpInputsListPagination, error) {
	queryParameters := &EncodingInputsTcpAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.TcpInputsListPagination
	err := api.apiClient.Get("/encoding/inputs/tcp", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsTcpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsTcpAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsTcpAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
