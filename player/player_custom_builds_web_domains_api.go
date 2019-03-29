package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type PlayerCustomBuildsWebDomainsApi struct {
    apiClient *common.ApiClient
}

func NewPlayerCustomBuildsWebDomainsApi(configs ...func(*common.ApiClient)) (*PlayerCustomBuildsWebDomainsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerCustomBuildsWebDomainsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerCustomBuildsWebDomainsApi) List() (*pagination.CustomWebPlayerBuildDomainsListPagination, error) {
    var resp *pagination.CustomWebPlayerBuildDomainsListPagination
    reqParams := func(params *common.RequestParams) {
	}
    err := api.apiClient.Get("/player/custom-builds/web/domains", &resp, reqParams)
    return resp, err
}
func (api *PlayerCustomBuildsWebDomainsApi) Get(domainId string) (*model.CustomWebPlayerBuildDomain, error) {
    var resp *model.CustomWebPlayerBuildDomain
    reqParams := func(params *common.RequestParams) {
        params.PathParams["domain_id"] = domainId
	}
    err := api.apiClient.Get("/player/custom-builds/web/domains/{domain_id}", &resp, reqParams)
    return resp, err
}
func (api *PlayerCustomBuildsWebDomainsApi) Create(customWebPlayerBuildDomain model.CustomWebPlayerBuildDomain) (*model.CustomWebPlayerBuildDomain, error) {
    payload := model.CustomWebPlayerBuildDomain(customWebPlayerBuildDomain)
    
    err := api.apiClient.Post("/player/custom-builds/web/domains", &payload)
    return &payload, err
}
func (api *PlayerCustomBuildsWebDomainsApi) Delete(domainId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["domain_id"] = domainId
	}
    err := api.apiClient.Delete("/player/custom-builds/web/domains/{domain_id}", &resp, reqParams)
    return resp, err
}
