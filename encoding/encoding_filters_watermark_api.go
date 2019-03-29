package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersWatermarkApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersWatermarkCustomdataApi
}

func NewEncodingFiltersWatermarkApi(configs ...func(*common.ApiClient)) (*EncodingFiltersWatermarkApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersWatermarkApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersWatermarkCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersWatermarkApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Delete("/encoding/filters/watermark/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersWatermarkApi) Get(filterId string) (*model.WatermarkFilter, error) {
    var resp *model.WatermarkFilter
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Get("/encoding/filters/watermark/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersWatermarkApi) List(queryParams ...func(*query.WatermarkFilterListQueryParams)) (*pagination.WatermarkFiltersListPagination, error) {
    queryParameters := &query.WatermarkFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.WatermarkFiltersListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/filters/watermark", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersWatermarkApi) Create(watermarkFilter model.WatermarkFilter) (*model.WatermarkFilter, error) {
    payload := model.WatermarkFilter(watermarkFilter)
    
    err := api.apiClient.Post("/encoding/filters/watermark", &payload)
    return &payload, err
}
