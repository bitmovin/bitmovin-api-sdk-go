package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioMp2API communicates with '/encoding/configurations/audio/mp2' endpoints
type EncodingConfigurationsAudioMp2API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/mp2/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioMp2CustomdataAPI
}

// NewEncodingConfigurationsAudioMp2API constructor for EncodingConfigurationsAudioMp2API that takes options as argument
func NewEncodingConfigurationsAudioMp2API(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioMp2API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioMp2APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioMp2APIWithClient constructor for EncodingConfigurationsAudioMp2API that takes an APIClient as argument
func NewEncodingConfigurationsAudioMp2APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioMp2API {
	a := &EncodingConfigurationsAudioMp2API{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioMp2CustomdataAPIWithClient(apiClient)

	return a
}

// Create MP2 Codec Configuration
func (api *EncodingConfigurationsAudioMp2API) Create(mp2AudioConfiguration model.Mp2AudioConfiguration) (*model.Mp2AudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Mp2AudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/mp2", &mp2AudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete MP2 Codec Configuration
func (api *EncodingConfigurationsAudioMp2API) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/mp2/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get MP2 Codec Configuration Details
func (api *EncodingConfigurationsAudioMp2API) Get(configurationId string) (*model.Mp2AudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.Mp2AudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/mp2/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List MP2 Configurations
func (api *EncodingConfigurationsAudioMp2API) List(queryParams ...func(*EncodingConfigurationsAudioMp2APIListQueryParams)) (*pagination.Mp2AudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioMp2APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.Mp2AudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/mp2", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioMp2APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioMp2APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioMp2APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
