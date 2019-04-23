package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type PlayerLicensesApi struct {
    apiClient *common.ApiClient
    Analytics *PlayerLicensesAnalyticsApi
    Domains *PlayerLicensesDomainsApi
    ThirdPartyLicensing *PlayerLicensesThirdPartyLicensingApi
}

func NewPlayerLicensesApi(configs ...func(*common.ApiClient)) (*PlayerLicensesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerLicensesApi{apiClient: apiClient}

    analyticsApi, err := NewPlayerLicensesAnalyticsApi(configs...)
    api.Analytics = analyticsApi
    domainsApi, err := NewPlayerLicensesDomainsApi(configs...)
    api.Domains = domainsApi
    thirdPartyLicensingApi, err := NewPlayerLicensesThirdPartyLicensingApi(configs...)
    api.ThirdPartyLicensing = thirdPartyLicensingApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerLicensesApi) Create(playerLicense model.PlayerLicense) (*model.PlayerLicense, error) {
    payload := model.PlayerLicense(playerLicense)
    
    err := api.apiClient.Post("/player/licenses", &payload)
    return &payload, err
}
func (api *PlayerLicensesApi) List(queryParams ...func(*query.PlayerLicenseListQueryParams)) (*pagination.PlayerLicensesListPagination, error) {
    queryParameters := &query.PlayerLicenseListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.PlayerLicensesListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/player/licenses", &resp, reqParams)
    return resp, err
}
func (api *PlayerLicensesApi) Get(licenseId string) (*model.PlayerLicense, error) {
    var resp *model.PlayerLicense
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
	}
    err := api.apiClient.Get("/player/licenses/{license_id}", &resp, reqParams)
    return resp, err
}
