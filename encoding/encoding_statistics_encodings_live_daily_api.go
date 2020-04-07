package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
    "time"
)

type EncodingStatisticsEncodingsLiveDailyApi struct {
    apiClient *common.ApiClient
}

func NewEncodingStatisticsEncodingsLiveDailyApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsLiveDailyApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsLiveDailyApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsLiveDailyApi) ListByDateRange(from time.Time, to time.Time, queryParams ...func(*query.EncodingStatisticsLiveListByDateRangeQueryParams)) (*pagination.EncodingStatisticsLivesListByDateRangePagination, error) {
    queryParameters := &query.EncodingStatisticsLiveListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["from"] = from
        params.PathParams["to"] = to
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EncodingStatisticsLivesListByDateRangePagination
    err := api.apiClient.Get("/encoding/statistics/encodings/live/daily/{from}/{to}", nil, &responseModel, reqParams)
    return responseModel, err
}

