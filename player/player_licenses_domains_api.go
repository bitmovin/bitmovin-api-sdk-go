package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type PlayerLicensesDomainsApi struct {
    apiClient *common.ApiClient
}

func NewPlayerLicensesDomainsApi(configs ...func(*common.ApiClient)) (*PlayerLicensesDomainsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerLicensesDomainsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerLicensesDomainsApi) Create(licenseId string, domain model.Domain) (*model.Domain, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }
    payload := model.Domain(domain)
    
    err := api.apiClient.Post("/player/licenses/{license_id}/domains", &payload, reqParams)
    return &payload, err
}
func (api *PlayerLicensesDomainsApi) List(licenseId string) (*pagination.DomainsListPagination, error) {
    var resp *pagination.DomainsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
	}
    err := api.apiClient.Get("/player/licenses/{license_id}/domains", &resp, reqParams)
    return resp, err
}
func (api *PlayerLicensesDomainsApi) Delete(licenseId string, domainId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
        params.PathParams["domain_id"] = domainId
	}
    err := api.apiClient.Delete("/player/licenses/{license_id}/domains/{domain_id}", &resp, reqParams)
    return resp, err
}
