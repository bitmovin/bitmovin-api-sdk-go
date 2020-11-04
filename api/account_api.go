package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AccountAPI intermediary API object with no endpoints
type AccountAPI struct {
	apiClient *apiclient.APIClient

	// Information communicates with '/account/information' endpoints
	Information *AccountInformationAPI
	// ApiKeys communicates with '/account/api-keys' endpoints
	ApiKeys *AccountApiKeysAPI
	// Organizations communicates with '/account/organizations' endpoints
	Organizations *AccountOrganizationsAPI
}

// NewAccountAPI constructor for AccountAPI that takes options as argument
func NewAccountAPI(options ...apiclient.APIClientOption) (*AccountAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAccountAPIWithClient(apiClient), nil
}

// NewAccountAPIWithClient constructor for AccountAPI that takes an APIClient as argument
func NewAccountAPIWithClient(apiClient *apiclient.APIClient) *AccountAPI {
	a := &AccountAPI{apiClient: apiClient}
	a.Information = NewAccountInformationAPIWithClient(apiClient)
	a.ApiKeys = NewAccountApiKeysAPIWithClient(apiClient)
	a.Organizations = NewAccountOrganizationsAPIWithClient(apiClient)

	return a
}
