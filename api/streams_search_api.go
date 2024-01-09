package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// StreamsSearchAPI communicates with '/streams/search' endpoints
type StreamsSearchAPI struct {
	apiClient *apiclient.APIClient
}

// NewStreamsSearchAPI constructor for StreamsSearchAPI that takes options as argument
func NewStreamsSearchAPI(options ...apiclient.APIClientOption) (*StreamsSearchAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsSearchAPIWithClient(apiClient), nil
}

// NewStreamsSearchAPIWithClient constructor for StreamsSearchAPI that takes an APIClient as argument
func NewStreamsSearchAPIWithClient(apiClient *apiclient.APIClient) *StreamsSearchAPI {
	a := &StreamsSearchAPI{apiClient: apiClient}
	return a
}

// List Get paginated search results of VOD and Live streams
func (api *StreamsSearchAPI) List(queryParams ...func(*StreamsSearchAPIListQueryParams)) (*pagination.StreamsResponsesListPagination, error) {
	queryParameters := &StreamsSearchAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StreamsResponsesListPagination
	err := api.apiClient.Get("/streams/search", nil, &responseModel, reqParams)
	return &responseModel, err
}

// StreamsSearchAPIListQueryParams contains all query parameters for the List endpoint
type StreamsSearchAPIListQueryParams struct {
	Offset           int32                    `query:"offset"`
	Limit            int32                    `query:"limit"`
	Query            string                   `query:"query"`
	Type_            model.StreamsType        `query:"type"`
	Status           model.StreamsVideoStatus `query:"status"`
	CreatedBefore    string                   `query:"createdBefore"`
	CreatedAfter     string                   `query:"createdAfter"`
	Signed           bool                     `query:"signed"`
	DomainRestricted bool                     `query:"domainRestricted"`
}

// Params will return a map of query parameters
func (q *StreamsSearchAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
