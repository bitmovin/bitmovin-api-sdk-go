package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioAc3API communicates with '/encoding/configurations/audio/ac3' endpoints
type EncodingConfigurationsAudioAc3API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/ac3/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioAc3CustomdataAPI
}

// NewEncodingConfigurationsAudioAc3API constructor for EncodingConfigurationsAudioAc3API that takes options as argument
func NewEncodingConfigurationsAudioAc3API(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioAc3API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioAc3APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioAc3APIWithClient constructor for EncodingConfigurationsAudioAc3API that takes an APIClient as argument
func NewEncodingConfigurationsAudioAc3APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioAc3API {
	a := &EncodingConfigurationsAudioAc3API{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioAc3CustomdataAPIWithClient(apiClient)

	return a
}

// Create AC3 Codec Configuration
func (api *EncodingConfigurationsAudioAc3API) Create(ac3AudioConfiguration model.Ac3AudioConfiguration) (*model.Ac3AudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Ac3AudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/ac3", &ac3AudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete AC3 Codec Configuration
func (api *EncodingConfigurationsAudioAc3API) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/ac3/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get AC3 Codec Configuration Details
func (api *EncodingConfigurationsAudioAc3API) Get(configurationId string) (*model.Ac3AudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.Ac3AudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/ac3/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List AC3 Configurations
func (api *EncodingConfigurationsAudioAc3API) List(queryParams ...func(*EncodingConfigurationsAudioAc3APIListQueryParams)) (*pagination.Ac3AudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioAc3APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.Ac3AudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/ac3", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioAc3APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioAc3APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioAc3APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
