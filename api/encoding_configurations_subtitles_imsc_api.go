package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsSubtitlesImscAPI communicates with '/encoding/configurations/subtitles/imsc' endpoints
type EncodingConfigurationsSubtitlesImscAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/subtitles/imsc/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsSubtitlesImscCustomdataAPI
}

// NewEncodingConfigurationsSubtitlesImscAPI constructor for EncodingConfigurationsSubtitlesImscAPI that takes options as argument
func NewEncodingConfigurationsSubtitlesImscAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsSubtitlesImscAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsSubtitlesImscAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsSubtitlesImscAPIWithClient constructor for EncodingConfigurationsSubtitlesImscAPI that takes an APIClient as argument
func NewEncodingConfigurationsSubtitlesImscAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsSubtitlesImscAPI {
	a := &EncodingConfigurationsSubtitlesImscAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsSubtitlesImscCustomdataAPIWithClient(apiClient)

	return a
}

// Create IMSC subtitle configuration
func (api *EncodingConfigurationsSubtitlesImscAPI) Create(imscConfiguration model.ImscConfiguration) (*model.ImscConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.ImscConfiguration
	err := api.apiClient.Post("/encoding/configurations/subtitles/imsc", &imscConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete IMSC subtitle configuration
func (api *EncodingConfigurationsSubtitlesImscAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/subtitles/imsc/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get IMSC subtitle configuration details
func (api *EncodingConfigurationsSubtitlesImscAPI) Get(configurationId string) (*model.ImscConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.ImscConfiguration
	err := api.apiClient.Get("/encoding/configurations/subtitles/imsc/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List IMSC subtitle configurations
func (api *EncodingConfigurationsSubtitlesImscAPI) List(queryParams ...func(*EncodingConfigurationsSubtitlesImscAPIListQueryParams)) (*pagination.ImscConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsSubtitlesImscAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ImscConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/subtitles/imsc", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsSubtitlesImscAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsSubtitlesImscAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsSubtitlesImscAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
