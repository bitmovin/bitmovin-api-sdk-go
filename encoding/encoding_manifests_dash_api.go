package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashApi struct {
    apiClient *common.ApiClient
    Default *EncodingManifestsDashDefaultApi
    Customdata *EncodingManifestsDashCustomdataApi
    Periods *EncodingManifestsDashPeriodsApi
}

func NewEncodingManifestsDashApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashApi{apiClient: apiClient}

    defaultApi, err := NewEncodingManifestsDashDefaultApi(configs...)
    api.Default = defaultApi
    customdataApi, err := NewEncodingManifestsDashCustomdataApi(configs...)
    api.Customdata = customdataApi
    periodsApi, err := NewEncodingManifestsDashPeriodsApi(configs...)
    api.Periods = periodsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashApi) Status(manifestId string) (*model.ModelTask, error) {
    var resp *model.ModelTask
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
	}
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/status", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsDashApi) Start(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    
    var payload model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/start", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsDashApi) List(queryParams ...func(*query.DashManifestListQueryParams)) (*pagination.DashManifestsListPagination, error) {
    queryParameters := &query.DashManifestListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.DashManifestsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/dash", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsDashApi) Create(dashManifest model.DashManifest) (*model.DashManifest, error) {
    payload := model.DashManifest(dashManifest)
    
    err := api.apiClient.Post("/encoding/manifests/dash", &payload)
    return &payload, err
}
func (api *EncodingManifestsDashApi) Get(manifestId string) (*model.DashManifest, error) {
    var resp *model.DashManifest
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
	}
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsDashApi) Stop(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    
    var payload model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/stop", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsDashApi) Delete(manifestId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
	}
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}", &resp, reqParams)
    return resp, err
}
