package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioDolbyDigitalAPI communicates with '/encoding/configurations/audio/dolby-digital' endpoints
type EncodingConfigurationsAudioDolbyDigitalAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/dolby-digital/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioDolbyDigitalCustomdataAPI
}

// NewEncodingConfigurationsAudioDolbyDigitalAPI constructor for EncodingConfigurationsAudioDolbyDigitalAPI that takes options as argument
func NewEncodingConfigurationsAudioDolbyDigitalAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDolbyDigitalAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioDolbyDigitalAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDolbyDigitalAPIWithClient constructor for EncodingConfigurationsAudioDolbyDigitalAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDolbyDigitalAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDolbyDigitalAPI {
	a := &EncodingConfigurationsAudioDolbyDigitalAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioDolbyDigitalCustomdataAPIWithClient(apiClient)

	return a
}

// Create Dolby Digital Codec Configuration
func (api *EncodingConfigurationsAudioDolbyDigitalAPI) Create(dolbyDigitalAudioConfiguration model.DolbyDigitalAudioConfiguration) (*model.DolbyDigitalAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.DolbyDigitalAudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/dolby-digital", &dolbyDigitalAudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Dolby Digital Codec Configuration
func (api *EncodingConfigurationsAudioDolbyDigitalAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/dolby-digital/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Dolby Digital Codec Configuration Details
func (api *EncodingConfigurationsAudioDolbyDigitalAPI) Get(configurationId string) (*model.DolbyDigitalAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.DolbyDigitalAudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/dolby-digital/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Dolby Digital Codec Configurations
func (api *EncodingConfigurationsAudioDolbyDigitalAPI) List(queryParams ...func(*EncodingConfigurationsAudioDolbyDigitalAPIListQueryParams)) (*pagination.DolbyDigitalAudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioDolbyDigitalAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DolbyDigitalAudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/dolby-digital", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioDolbyDigitalAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioDolbyDigitalAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioDolbyDigitalAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
