package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsLicensesDomainsApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsLicensesDomainsApi(configs ...func(*common.ApiClient)) (*AnalyticsLicensesDomainsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsLicensesDomainsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsLicensesDomainsApi) Delete(licenseId string, domainId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
        params.PathParams["domain_id"] = domainId
	}
    err := api.apiClient.Delete("/analytics/licenses/{license_id}/domains/{domain_id}", &resp, reqParams)
    return resp, err
}
func (api *AnalyticsLicensesDomainsApi) Get(licenseId string) (*model.DomainList, error) {
    var resp *model.DomainList
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
	}
    err := api.apiClient.Get("/analytics/licenses/{license_id}/domains", &resp, reqParams)
    return resp, err
}
func (api *AnalyticsLicensesDomainsApi) Create(licenseId string, domain model.Domain) (*model.Domain, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }
    payload := model.Domain(domain)
    
    err := api.apiClient.Post("/analytics/licenses/{license_id}/domains", &payload, reqParams)
    return &payload, err
}
