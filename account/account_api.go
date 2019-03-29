package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type AccountApi struct {
    apiClient *common.ApiClient
    Information *AccountInformationApi
    Login *AccountLoginApi
    ApiKeys *AccountApiKeysApi
    Organizations *AccountOrganizationsApi
}

func NewAccountApi(configs ...func(*common.ApiClient)) (*AccountApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountApi{apiClient: apiClient}

    informationApi, err := NewAccountInformationApi(configs...)
    api.Information = informationApi
    loginApi, err := NewAccountLoginApi(configs...)
    api.Login = loginApi
    apiKeysApi, err := NewAccountApiKeysApi(configs...)
    api.ApiKeys = apiKeysApi
    organizationsApi, err := NewAccountOrganizationsApi(configs...)
    api.Organizations = organizationsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

