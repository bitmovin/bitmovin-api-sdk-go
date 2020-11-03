package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AccountInformationAPI communicates with '/account/information' endpoints
type AccountInformationAPI struct {
    apiClient *apiclient.APIClient
}

// NewAccountInformationAPI constructor for AccountInformationAPI that takes options as argument
func NewAccountInformationAPI(options ...apiclient.APIClientOption) (*AccountInformationAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAccountInformationAPIWithClient(apiClient), nil
}

// NewAccountInformationAPIWithClient constructor for AccountInformationAPI that takes an APIClient as argument
func NewAccountInformationAPIWithClient(apiClient *apiclient.APIClient) *AccountInformationAPI {
    a := &AccountInformationAPI{apiClient: apiClient}
    return a
}

// Get Current Account Information
func (api *AccountInformationAPI) Get() (*model.AccountInformation, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AccountInformation
    err := api.apiClient.Get("/account/information", nil, &responseModel, reqParams)
    return &responseModel, err
}

