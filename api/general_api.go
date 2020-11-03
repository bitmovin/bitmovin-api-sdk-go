package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// GeneralAPI intermediary API object with no endpoints
type GeneralAPI struct {
    apiClient *apiclient.APIClient

    // ErrorDefinitions communicates with '/general/error-definitions' endpoints
    ErrorDefinitions *GeneralErrorDefinitionsAPI

}

// NewGeneralAPI constructor for GeneralAPI that takes options as argument
func NewGeneralAPI(options ...apiclient.APIClientOption) (*GeneralAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewGeneralAPIWithClient(apiClient), nil
}

// NewGeneralAPIWithClient constructor for GeneralAPI that takes an APIClient as argument
func NewGeneralAPIWithClient(apiClient *apiclient.APIClient) *GeneralAPI {
    a := &GeneralAPI{apiClient: apiClient}
    a.ErrorDefinitions = NewGeneralErrorDefinitionsAPIWithClient(apiClient)

    return a
}


