package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioMp3API communicates with '/encoding/configurations/audio/mp3' endpoints
type EncodingConfigurationsAudioMp3API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/mp3/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioMp3CustomdataAPI
}

// NewEncodingConfigurationsAudioMp3API constructor for EncodingConfigurationsAudioMp3API that takes options as argument
func NewEncodingConfigurationsAudioMp3API(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioMp3API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioMp3APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioMp3APIWithClient constructor for EncodingConfigurationsAudioMp3API that takes an APIClient as argument
func NewEncodingConfigurationsAudioMp3APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioMp3API {
	a := &EncodingConfigurationsAudioMp3API{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioMp3CustomdataAPIWithClient(apiClient)

	return a
}

// Create MP3 Codec Configuration
func (api *EncodingConfigurationsAudioMp3API) Create(mp3AudioConfiguration model.Mp3AudioConfiguration) (*model.Mp3AudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Mp3AudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/mp3", &mp3AudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete MP3 Codec Configuration
func (api *EncodingConfigurationsAudioMp3API) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/mp3/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get MP3 Codec Configuration Details
func (api *EncodingConfigurationsAudioMp3API) Get(configurationId string) (*model.Mp3AudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.Mp3AudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/mp3/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List MP3 Configurations
func (api *EncodingConfigurationsAudioMp3API) List(queryParams ...func(*EncodingConfigurationsAudioMp3APIListQueryParams)) (*pagination.Mp3AudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioMp3APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.Mp3AudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/mp3", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioMp3APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioMp3APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioMp3APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
