package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type PlayerCustomBuildsWebApi struct {
    apiClient *common.ApiClient
    Domains *PlayerCustomBuildsWebDomainsApi
    Status *PlayerCustomBuildsWebStatusApi
    Download *PlayerCustomBuildsWebDownloadApi
}

func NewPlayerCustomBuildsWebApi(configs ...func(*common.ApiClient)) (*PlayerCustomBuildsWebApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerCustomBuildsWebApi{apiClient: apiClient}

    domainsApi, err := NewPlayerCustomBuildsWebDomainsApi(configs...)
    api.Domains = domainsApi
    statusApi, err := NewPlayerCustomBuildsWebStatusApi(configs...)
    api.Status = statusApi
    downloadApi, err := NewPlayerCustomBuildsWebDownloadApi(configs...)
    api.Download = downloadApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerCustomBuildsWebApi) List() (*pagination.CustomPlayerBuildDetailssListPagination, error) {
    var resp *pagination.CustomPlayerBuildDetailssListPagination
    reqParams := func(params *common.RequestParams) {
	}
    err := api.apiClient.Get("/player/custom-builds/web", &resp, reqParams)
    return resp, err
}
func (api *PlayerCustomBuildsWebApi) Create(customPlayerBuildDetails model.CustomPlayerBuildDetails) (*model.CustomPlayerBuildDetails, error) {
    payload := model.CustomPlayerBuildDetails(customPlayerBuildDetails)
    
    err := api.apiClient.Post("/player/custom-builds/web", &payload)
    return &payload, err
}
func (api *PlayerCustomBuildsWebApi) Start(customBuildId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
    }
    
    var payload model.BitmovinResponse
    err := api.apiClient.Post("/player/custom-builds/web/{custom_build_id}/start", &payload, reqParams)
    return &payload, err
}
func (api *PlayerCustomBuildsWebApi) Get(customBuildId string) (*model.CustomPlayerBuildStatus, error) {
    var resp *model.CustomPlayerBuildStatus
    reqParams := func(params *common.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
	}
    err := api.apiClient.Get("/player/custom-builds/web/{custom_build_id}", &resp, reqParams)
    return resp, err
}
