package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersDeinterlaceApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersDeinterlaceCustomdataApi
}

func NewEncodingFiltersDeinterlaceApi(configs ...func(*common.ApiClient)) (*EncodingFiltersDeinterlaceApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersDeinterlaceApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersDeinterlaceCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersDeinterlaceApi) Create(deinterlaceFilter model.DeinterlaceFilter) (*model.DeinterlaceFilter, error) {
    payload := model.DeinterlaceFilter(deinterlaceFilter)
    
    err := api.apiClient.Post("/encoding/filters/deinterlace", &payload)
    return &payload, err
}
func (api *EncodingFiltersDeinterlaceApi) List(queryParams ...func(*query.DeinterlaceFilterListQueryParams)) (*pagination.DeinterlaceFiltersListPagination, error) {
    queryParameters := &query.DeinterlaceFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.DeinterlaceFiltersListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/filters/deinterlace", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersDeinterlaceApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Delete("/encoding/filters/deinterlace/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersDeinterlaceApi) Get(filterId string) (*model.DeinterlaceFilter, error) {
    var resp *model.DeinterlaceFilter
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Get("/encoding/filters/deinterlace/{filter_id}", &resp, reqParams)
    return resp, err
}
