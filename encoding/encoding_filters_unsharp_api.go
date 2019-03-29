package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersUnsharpApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersUnsharpCustomdataApi
}

func NewEncodingFiltersUnsharpApi(configs ...func(*common.ApiClient)) (*EncodingFiltersUnsharpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersUnsharpApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersUnsharpCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersUnsharpApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Delete("/encoding/filters/unsharp/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersUnsharpApi) List(queryParams ...func(*query.UnsharpFilterListQueryParams)) (*pagination.UnsharpFiltersListPagination, error) {
    queryParameters := &query.UnsharpFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.UnsharpFiltersListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/filters/unsharp", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersUnsharpApi) Get(filterId string) (*model.UnsharpFilter, error) {
    var resp *model.UnsharpFilter
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Get("/encoding/filters/unsharp/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersUnsharpApi) Create(unsharpFilter model.UnsharpFilter) (*model.UnsharpFilter, error) {
    payload := model.UnsharpFilter(unsharpFilter)
    
    err := api.apiClient.Post("/encoding/filters/unsharp", &payload)
    return &payload, err
}
