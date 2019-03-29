package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersCropApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersCropCustomdataApi
}

func NewEncodingFiltersCropApi(configs ...func(*common.ApiClient)) (*EncodingFiltersCropApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersCropApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersCropCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersCropApi) Create(cropFilter model.CropFilter) (*model.CropFilter, error) {
    payload := model.CropFilter(cropFilter)
    
    err := api.apiClient.Post("/encoding/filters/crop", &payload)
    return &payload, err
}
func (api *EncodingFiltersCropApi) List(queryParams ...func(*query.CropFilterListQueryParams)) (*pagination.CropFiltersListPagination, error) {
    queryParameters := &query.CropFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.CropFiltersListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/filters/crop", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersCropApi) Get(filterId string) (*model.CropFilter, error) {
    var resp *model.CropFilter
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Get("/encoding/filters/crop/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersCropApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Delete("/encoding/filters/crop/{filter_id}", &resp, reqParams)
    return resp, err
}
