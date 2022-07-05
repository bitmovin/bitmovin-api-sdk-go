package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingStatisticsCdnUsageAPI communicates with '/encoding/statistics/cdn/usage' endpoints
type EncodingStatisticsCdnUsageAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingStatisticsCdnUsageAPI constructor for EncodingStatisticsCdnUsageAPI that takes options as argument
func NewEncodingStatisticsCdnUsageAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsCdnUsageAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsCdnUsageAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsCdnUsageAPIWithClient constructor for EncodingStatisticsCdnUsageAPI that takes an APIClient as argument
func NewEncodingStatisticsCdnUsageAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsCdnUsageAPI {
	a := &EncodingStatisticsCdnUsageAPI{apiClient: apiClient}
	return a
}

// Get List CDN usage statistics within specific dates.
func (api *EncodingStatisticsCdnUsageAPI) Get(queryParams ...func(*EncodingStatisticsCdnUsageAPIGetQueryParams)) (*model.CdnUsageStatistics, error) {
	queryParameters := &EncodingStatisticsCdnUsageAPIGetQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel model.CdnUsageStatistics
	err := api.apiClient.Get("/encoding/statistics/cdn/usage", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingStatisticsCdnUsageAPIGetQueryParams contains all query parameters for the Get endpoint
type EncodingStatisticsCdnUsageAPIGetQueryParams struct {
	From model.DateTime `query:"from"`
	To   model.DateTime `query:"to"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsCdnUsageAPIGetQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
