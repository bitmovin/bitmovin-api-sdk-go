package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioDtsAPI communicates with '/encoding/configurations/audio/dts' endpoints
type EncodingConfigurationsAudioDtsAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/dts/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioDtsCustomdataAPI
}

// NewEncodingConfigurationsAudioDtsAPI constructor for EncodingConfigurationsAudioDtsAPI that takes options as argument
func NewEncodingConfigurationsAudioDtsAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDtsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioDtsAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDtsAPIWithClient constructor for EncodingConfigurationsAudioDtsAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDtsAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDtsAPI {
	a := &EncodingConfigurationsAudioDtsAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioDtsCustomdataAPIWithClient(apiClient)

	return a
}

// Create DTS Codec Configuration
func (api *EncodingConfigurationsAudioDtsAPI) Create(dtsAudioConfiguration model.DtsAudioConfiguration) (*model.DtsAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.DtsAudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/dts", &dtsAudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete DTS Codec Configuration
func (api *EncodingConfigurationsAudioDtsAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/dts/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get DTS Codec Configuration Details
func (api *EncodingConfigurationsAudioDtsAPI) Get(configurationId string) (*model.DtsAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.DtsAudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/dts/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List DTS Codec Configurations
func (api *EncodingConfigurationsAudioDtsAPI) List(queryParams ...func(*EncodingConfigurationsAudioDtsAPIListQueryParams)) (*pagination.DtsAudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioDtsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DtsAudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/dts", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioDtsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioDtsAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioDtsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
