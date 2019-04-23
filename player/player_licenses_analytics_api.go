package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type PlayerLicensesAnalyticsApi struct {
    apiClient *common.ApiClient
}

func NewPlayerLicensesAnalyticsApi(configs ...func(*common.ApiClient)) (*PlayerLicensesAnalyticsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerLicensesAnalyticsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerLicensesAnalyticsApi) Create(licenseId string, playerLicenseAnalytics model.PlayerLicenseAnalytics) (*model.PlayerLicenseAnalytics, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }
    payload := model.PlayerLicenseAnalytics(playerLicenseAnalytics)
    
    err := api.apiClient.Post("/player/licenses/{license_id}/analytics", &payload, reqParams)
    return &payload, err
}
func (api *PlayerLicensesAnalyticsApi) Delete(licenseId string) (*model.PlayerLicenseAnalytics, error) {
    var resp *model.PlayerLicenseAnalytics
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
	}
    err := api.apiClient.Delete("/player/licenses/{license_id}/analytics", &resp, reqParams)
    return resp, err
}
