package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountLimitsApi struct {
    apiClient *common.ApiClient
}

func NewAccountLimitsApi(configs ...func(*common.ApiClient)) (*AccountLimitsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountLimitsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountLimitsApi) List() (*pagination.ResourceLimitContainersListPagination, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *pagination.ResourceLimitContainersListPagination
    err := api.apiClient.Get("/account/limits", nil, &responseModel, reqParams)
    return responseModel, err
}

