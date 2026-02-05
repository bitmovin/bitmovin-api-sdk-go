package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// StreamsConfigDomainRestrictionAPI communicates with '/streams/config/domain-restriction' endpoints
type StreamsConfigDomainRestrictionAPI struct {
	apiClient *apiclient.APIClient
}

// NewStreamsConfigDomainRestrictionAPI constructor for StreamsConfigDomainRestrictionAPI that takes options as argument
func NewStreamsConfigDomainRestrictionAPI(options ...apiclient.APIClientOption) (*StreamsConfigDomainRestrictionAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsConfigDomainRestrictionAPIWithClient(apiClient), nil
}

// NewStreamsConfigDomainRestrictionAPIWithClient constructor for StreamsConfigDomainRestrictionAPI that takes an APIClient as argument
func NewStreamsConfigDomainRestrictionAPIWithClient(apiClient *apiclient.APIClient) *StreamsConfigDomainRestrictionAPI {
	a := &StreamsConfigDomainRestrictionAPI{apiClient: apiClient}
	return a
}

// PatchStreamsDomainRestriction Partially update streams domain restriction by id
func (api *StreamsConfigDomainRestrictionAPI) PatchStreamsDomainRestriction(domainRestrictionId string, streamsDomainRestrictionUpdateRequest model.StreamsDomainRestrictionUpdateRequest) (*model.StreamsDomainRestrictionResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["domain_restriction_id"] = domainRestrictionId
	}

	var responseModel model.StreamsDomainRestrictionResponse
	err := api.apiClient.Patch("/streams/config/domain-restriction/{domain_restriction_id}", &streamsDomainRestrictionUpdateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Create new streams domain restriction
func (api *StreamsConfigDomainRestrictionAPI) Create(streamsDomainRestrictionCreateRequest model.StreamsDomainRestrictionCreateRequest) (*model.StreamsDomainRestrictionResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.StreamsDomainRestrictionResponse
	err := api.apiClient.Post("/streams/config/domain-restriction", &streamsDomainRestrictionCreateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Delete streams domain restriction by id
func (api *StreamsConfigDomainRestrictionAPI) Delete(domainRestrictionId string) error {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["domain_restriction_id"] = domainRestrictionId
	}

	err := api.apiClient.Delete("/streams/config/domain-restriction/{domain_restriction_id}", nil, nil, reqParams)
	return err
}

// Get streams domain restriction config by id
func (api *StreamsConfigDomainRestrictionAPI) Get(domainRestrictionId string) (*model.StreamsDomainRestrictionResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["domain_restriction_id"] = domainRestrictionId
	}

	var responseModel model.StreamsDomainRestrictionResponse
	err := api.apiClient.Get("/streams/config/domain-restriction/{domain_restriction_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Get paginated list of domain restriction configurations
func (api *StreamsConfigDomainRestrictionAPI) List(queryParams ...func(*StreamsConfigDomainRestrictionAPIListQueryParams)) (*pagination.StreamsDomainRestrictionResponsesListPagination, error) {
	queryParameters := &StreamsConfigDomainRestrictionAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StreamsDomainRestrictionResponsesListPagination
	err := api.apiClient.Get("/streams/config/domain-restriction", nil, &responseModel, reqParams)
	return &responseModel, err
}

// StreamsConfigDomainRestrictionAPIListQueryParams contains all query parameters for the List endpoint
type StreamsConfigDomainRestrictionAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *StreamsConfigDomainRestrictionAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
