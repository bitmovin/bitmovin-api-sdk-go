package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// PlayerLicensesThirdPartyLicensingAPI communicates with '/player/licenses/{license_id}/third-party-licensing' endpoints
type PlayerLicensesThirdPartyLicensingAPI struct {
	apiClient *apiclient.APIClient
}

// NewPlayerLicensesThirdPartyLicensingAPI constructor for PlayerLicensesThirdPartyLicensingAPI that takes options as argument
func NewPlayerLicensesThirdPartyLicensingAPI(options ...apiclient.APIClientOption) (*PlayerLicensesThirdPartyLicensingAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewPlayerLicensesThirdPartyLicensingAPIWithClient(apiClient), nil
}

// NewPlayerLicensesThirdPartyLicensingAPIWithClient constructor for PlayerLicensesThirdPartyLicensingAPI that takes an APIClient as argument
func NewPlayerLicensesThirdPartyLicensingAPIWithClient(apiClient *apiclient.APIClient) *PlayerLicensesThirdPartyLicensingAPI {
	a := &PlayerLicensesThirdPartyLicensingAPI{apiClient: apiClient}
	return a
}

// Create Enable Third Party Licensing
func (api *PlayerLicensesThirdPartyLicensingAPI) Create(licenseId string, playerThirdPartyLicensing model.PlayerThirdPartyLicensing) (*model.PlayerThirdPartyLicensing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_id"] = licenseId
	}

	var responseModel model.PlayerThirdPartyLicensing
	err := api.apiClient.Post("/player/licenses/{license_id}/third-party-licensing", &playerThirdPartyLicensing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Third Party Licensing Configuration
func (api *PlayerLicensesThirdPartyLicensingAPI) Delete(licenseId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_id"] = licenseId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/player/licenses/{license_id}/third-party-licensing", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Third Party Licensing Configuration
func (api *PlayerLicensesThirdPartyLicensingAPI) Get(licenseId string) (*model.PlayerThirdPartyLicensing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_id"] = licenseId
	}

	var responseModel model.PlayerThirdPartyLicensing
	err := api.apiClient.Get("/player/licenses/{license_id}/third-party-licensing", nil, &responseModel, reqParams)
	return &responseModel, err
}
