package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsApi struct {
    apiClient *common.ApiClient
    CustomXmlElements *EncodingManifestsDashPeriodsCustomXmlElementsApi
    Adaptationsets *EncodingManifestsDashPeriodsAdaptationsetsApi
}

func NewEncodingManifestsDashPeriodsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsApi{apiClient: apiClient}

    customXmlElementsApi, err := NewEncodingManifestsDashPeriodsCustomXmlElementsApi(configs...)
    api.CustomXmlElements = customXmlElementsApi
    adaptationsetsApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsApi(configs...)
    api.Adaptationsets = adaptationsetsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsApi) Get(manifestId string, periodId string) (*model.Period, error) {
    var resp *model.Period
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
	}
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsDashPeriodsApi) Create(manifestId string, period model.Period) (*model.Period, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    payload := model.Period(period)
    
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsDashPeriodsApi) List(manifestId string, queryParams ...func(*query.PeriodListQueryParams)) (*pagination.PeriodsListPagination, error) {
    queryParameters := &query.PeriodListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.PeriodsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsDashPeriodsApi) Delete(manifestId string, periodId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
	}
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}", &resp, reqParams)
    return resp, err
}
