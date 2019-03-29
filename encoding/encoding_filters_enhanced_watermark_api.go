package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersEnhancedWatermarkApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersEnhancedWatermarkCustomdataApi
}

func NewEncodingFiltersEnhancedWatermarkApi(configs ...func(*common.ApiClient)) (*EncodingFiltersEnhancedWatermarkApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersEnhancedWatermarkApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersEnhancedWatermarkCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersEnhancedWatermarkApi) Create(enhancedWatermarkFilter model.EnhancedWatermarkFilter) (*model.EnhancedWatermarkFilter, error) {
    payload := model.EnhancedWatermarkFilter(enhancedWatermarkFilter)
    
    err := api.apiClient.Post("/encoding/filters/enhanced-watermark", &payload)
    return &payload, err
}
func (api *EncodingFiltersEnhancedWatermarkApi) Get(filterId string) (*model.EnhancedWatermarkFilter, error) {
    var resp *model.EnhancedWatermarkFilter
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Get("/encoding/filters/enhanced-watermark/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersEnhancedWatermarkApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Delete("/encoding/filters/enhanced-watermark/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersEnhancedWatermarkApi) List(queryParams ...func(*query.EnhancedWatermarkFilterListQueryParams)) (*pagination.EnhancedWatermarkFiltersListPagination, error) {
    queryParameters := &query.EnhancedWatermarkFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.EnhancedWatermarkFiltersListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/filters/enhanced-watermark", &resp, reqParams)
    return resp, err
}
