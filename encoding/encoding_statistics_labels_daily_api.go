package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
    "time"
)

type EncodingStatisticsLabelsDailyApi struct {
    apiClient *common.ApiClient
}

func NewEncodingStatisticsLabelsDailyApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsLabelsDailyApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsLabelsDailyApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsLabelsDailyApi) ListByDateRange(from time.Time, to time.Time, queryParams ...func(*query.DailyStatisticsPerLabelListByDateRangeQueryParams)) (*pagination.DailyStatisticsPerLabelsListByDateRangePagination, error) {
    queryParameters := &query.DailyStatisticsPerLabelListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.DailyStatisticsPerLabelsListByDateRangePagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["from"] = from
        params.PathParams["to"] = to
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/statistics/labels/daily/{from}/{to}", &resp, reqParams)
    return resp, err
}
func (api *EncodingStatisticsLabelsDailyApi) List(queryParams ...func(*query.DailyStatisticsPerLabelListQueryParams)) (*pagination.DailyStatisticsPerLabelsListPagination, error) {
    queryParameters := &query.DailyStatisticsPerLabelListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.DailyStatisticsPerLabelsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/statistics/labels/daily", &resp, reqParams)
    return resp, err
}
