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

func (api *AccountLoginApi) Create(login model.Login) (*model.AccountInformation, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AccountInformation
    err := api.apiClient.Post("/account/login", &login, &responseModel, reqParams)
    return responseModel, err
}

