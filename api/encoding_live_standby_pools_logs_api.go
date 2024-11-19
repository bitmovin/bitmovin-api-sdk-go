package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingLiveStandbyPoolsLogsAPI communicates with '/encoding/live/standby-pools/{pool_id}/logs' endpoints
type EncodingLiveStandbyPoolsLogsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingLiveStandbyPoolsLogsAPI constructor for EncodingLiveStandbyPoolsLogsAPI that takes options as argument
func NewEncodingLiveStandbyPoolsLogsAPI(options ...apiclient.APIClientOption) (*EncodingLiveStandbyPoolsLogsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveStandbyPoolsLogsAPIWithClient(apiClient), nil
}

// NewEncodingLiveStandbyPoolsLogsAPIWithClient constructor for EncodingLiveStandbyPoolsLogsAPI that takes an APIClient as argument
func NewEncodingLiveStandbyPoolsLogsAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveStandbyPoolsLogsAPI {
	a := &EncodingLiveStandbyPoolsLogsAPI{apiClient: apiClient}
	return a
}

// List event logs for a standby pool
func (api *EncodingLiveStandbyPoolsLogsAPI) List(poolId string, queryParams ...func(*EncodingLiveStandbyPoolsLogsAPIListQueryParams)) (*pagination.LiveStandbyPoolEventLogsListPagination, error) {
	queryParameters := &EncodingLiveStandbyPoolsLogsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["pool_id"] = poolId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.LiveStandbyPoolEventLogsListPagination
	err := api.apiClient.Get("/encoding/live/standby-pools/{pool_id}/logs", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingLiveStandbyPoolsLogsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingLiveStandbyPoolsLogsAPIListQueryParams struct {
	Offset    int32  `query:"offset"`
	Limit     int32  `query:"limit"`
	Sort      string `query:"sort"`
	EventType string `query:"eventType"`
}

// Params will return a map of query parameters
func (q *EncodingLiveStandbyPoolsLogsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
