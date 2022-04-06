package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsCdnAPI communicates with '/encoding/outputs/cdn' endpoints
type EncodingOutputsCdnAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingOutputsCdnAPI constructor for EncodingOutputsCdnAPI that takes options as argument
func NewEncodingOutputsCdnAPI(options ...apiclient.APIClientOption) (*EncodingOutputsCdnAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsCdnAPIWithClient(apiClient), nil
}

// NewEncodingOutputsCdnAPIWithClient constructor for EncodingOutputsCdnAPI that takes an APIClient as argument
func NewEncodingOutputsCdnAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsCdnAPI {
	a := &EncodingOutputsCdnAPI{apiClient: apiClient}
	return a
}

// Get CDN Output Details
func (api *EncodingOutputsCdnAPI) Get(outputId string) (*model.CdnOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CdnOutput
	err := api.apiClient.Get("/encoding/outputs/cdn/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List CDN Outputs
func (api *EncodingOutputsCdnAPI) List(queryParams ...func(*EncodingOutputsCdnAPIListQueryParams)) (*pagination.CdnOutputsListPagination, error) {
	queryParameters := &EncodingOutputsCdnAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.CdnOutputsListPagination
	err := api.apiClient.Get("/encoding/outputs/cdn", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsCdnAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsCdnAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsCdnAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
