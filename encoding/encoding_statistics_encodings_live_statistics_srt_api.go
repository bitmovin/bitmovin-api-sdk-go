package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingStatisticsEncodingsLiveStatisticsSrtApi struct {
    apiClient *common.ApiClient
}

func NewEncodingStatisticsEncodingsLiveStatisticsSrtApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsLiveStatisticsSrtApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsLiveStatisticsSrtApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsLiveStatisticsSrtApi) List(encodingId string, queryParams ...func(*query.SrtStatisticsListQueryParams)) (*pagination.SrtStatisticssListPagination, error) {
    queryParameters := &query.SrtStatisticsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SrtStatisticssListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics/srt", &resp, reqParams)
    return resp, err
}
