package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AnalyticsLicensesApi struct {
    apiClient *common.ApiClient
    Domains *AnalyticsLicensesDomainsApi
}

func NewAnalyticsLicensesApi(configs ...func(*common.ApiClient)) (*AnalyticsLicensesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsLicensesApi{apiClient: apiClient}

    domainsApi, err := NewAnalyticsLicensesDomainsApi(configs...)
    api.Domains = domainsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsLicensesApi) Create(analyticsLicense model.AnalyticsLicense) (*model.AnalyticsLicense, error) {
    payload := model.AnalyticsLicense(analyticsLicense)
    
    err := api.apiClient.Post("/analytics/licenses", &payload)
    return &payload, err
}
func (api *AnalyticsLicensesApi) Get(licenseId string) (*model.AnalyticsLicense, error) {
    var resp *model.AnalyticsLicense
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
	}
    err := api.apiClient.Get("/analytics/licenses/{license_id}", &resp, reqParams)
    return resp, err
}
func (api *AnalyticsLicensesApi) List() (*pagination.AnalyticsLicensesListPagination, error) {
    var resp *pagination.AnalyticsLicensesListPagination
    reqParams := func(params *common.RequestParams) {
	}
    err := api.apiClient.Get("/analytics/licenses", &resp, reqParams)
    return resp, err
}
