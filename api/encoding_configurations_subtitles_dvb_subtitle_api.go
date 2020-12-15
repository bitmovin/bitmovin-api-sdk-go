package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsSubtitlesDvbSubtitleAPI communicates with '/encoding/configurations/subtitles/dvb-subtitle' endpoints
type EncodingConfigurationsSubtitlesDvbSubtitleAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/subtitles/dvb-subtitle/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI
}

// NewEncodingConfigurationsSubtitlesDvbSubtitleAPI constructor for EncodingConfigurationsSubtitlesDvbSubtitleAPI that takes options as argument
func NewEncodingConfigurationsSubtitlesDvbSubtitleAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsSubtitlesDvbSubtitleAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsSubtitlesDvbSubtitleAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsSubtitlesDvbSubtitleAPIWithClient constructor for EncodingConfigurationsSubtitlesDvbSubtitleAPI that takes an APIClient as argument
func NewEncodingConfigurationsSubtitlesDvbSubtitleAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsSubtitlesDvbSubtitleAPI {
	a := &EncodingConfigurationsSubtitlesDvbSubtitleAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPIWithClient(apiClient)

	return a
}

// Create DVB-SUB subtitle configuration
func (api *EncodingConfigurationsSubtitlesDvbSubtitleAPI) Create(dvbSubtitleConfiguration model.DvbSubtitleConfiguration) (*model.DvbSubtitleConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.DvbSubtitleConfiguration
	err := api.apiClient.Post("/encoding/configurations/subtitles/dvb-subtitle", &dvbSubtitleConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete DVB-SUB subtitle configuration
func (api *EncodingConfigurationsSubtitlesDvbSubtitleAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/subtitles/dvb-subtitle/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get DVB-SUB subtitle configuration details
func (api *EncodingConfigurationsSubtitlesDvbSubtitleAPI) Get(configurationId string) (*model.DvbSubtitleConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.DvbSubtitleConfiguration
	err := api.apiClient.Get("/encoding/configurations/subtitles/dvb-subtitle/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List DVB-SUB subtitle configurations
func (api *EncodingConfigurationsSubtitlesDvbSubtitleAPI) List(queryParams ...func(*EncodingConfigurationsSubtitlesDvbSubtitleAPIListQueryParams)) (*pagination.DvbSubtitleConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsSubtitlesDvbSubtitleAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DvbSubtitleConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/subtitles/dvb-subtitle", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsSubtitlesDvbSubtitleAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsSubtitlesDvbSubtitleAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsSubtitlesDvbSubtitleAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
