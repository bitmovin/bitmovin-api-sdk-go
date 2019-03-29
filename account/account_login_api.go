package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AccountLoginApi struct {
    apiClient *common.ApiClient
}

func NewAccountLoginApi(configs ...func(*common.ApiClient)) (*AccountLoginApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountLoginApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountLoginApi) Create(login model.AccountInformation) (*model.AccountInformation, error) {
    payload := model.AccountInformation(login)
    
    err := api.apiClient.Post("/account/login", &payload)
    return &payload, err
}
