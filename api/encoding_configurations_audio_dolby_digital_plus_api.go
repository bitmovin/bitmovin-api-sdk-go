package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioDolbyDigitalPlusAPI communicates with '/encoding/configurations/audio/dolby-digital-plus' endpoints
type EncodingConfigurationsAudioDolbyDigitalPlusAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/dolby-digital-plus/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI
}

// NewEncodingConfigurationsAudioDolbyDigitalPlusAPI constructor for EncodingConfigurationsAudioDolbyDigitalPlusAPI that takes options as argument
func NewEncodingConfigurationsAudioDolbyDigitalPlusAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDolbyDigitalPlusAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioDolbyDigitalPlusAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDolbyDigitalPlusAPIWithClient constructor for EncodingConfigurationsAudioDolbyDigitalPlusAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDolbyDigitalPlusAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDolbyDigitalPlusAPI {
	a := &EncodingConfigurationsAudioDolbyDigitalPlusAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPIWithClient(apiClient)

	return a
}

// Create Dolby Digital Plus Codec Configuration
func (api *EncodingConfigurationsAudioDolbyDigitalPlusAPI) Create(dolbyDigitalPlusAudioConfiguration model.DolbyDigitalPlusAudioConfiguration) (*model.DolbyDigitalPlusAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.DolbyDigitalPlusAudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/dolby-digital-plus", &dolbyDigitalPlusAudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Dolby Digital Plus Codec Configuration
func (api *EncodingConfigurationsAudioDolbyDigitalPlusAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/dolby-digital-plus/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Dolby Digital Plus Codec Configuration Details
func (api *EncodingConfigurationsAudioDolbyDigitalPlusAPI) Get(configurationId string) (*model.DolbyDigitalPlusAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.DolbyDigitalPlusAudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/dolby-digital-plus/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Dolby Digital Plus Codec Configurations
func (api *EncodingConfigurationsAudioDolbyDigitalPlusAPI) List(queryParams ...func(*EncodingConfigurationsAudioDolbyDigitalPlusAPIListQueryParams)) (*pagination.DolbyDigitalPlusAudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioDolbyDigitalPlusAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DolbyDigitalPlusAudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/dolby-digital-plus", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioDolbyDigitalPlusAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioDolbyDigitalPlusAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioDolbyDigitalPlusAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
