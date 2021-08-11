package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioDtsxAPI communicates with '/encoding/configurations/audio/dtsx' endpoints
type EncodingConfigurationsAudioDtsxAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/dtsx/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioDtsxCustomdataAPI
}

// NewEncodingConfigurationsAudioDtsxAPI constructor for EncodingConfigurationsAudioDtsxAPI that takes options as argument
func NewEncodingConfigurationsAudioDtsxAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDtsxAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioDtsxAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDtsxAPIWithClient constructor for EncodingConfigurationsAudioDtsxAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDtsxAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDtsxAPI {
	a := &EncodingConfigurationsAudioDtsxAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioDtsxCustomdataAPIWithClient(apiClient)

	return a
}

// Create DTS:X Codec Configuration
func (api *EncodingConfigurationsAudioDtsxAPI) Create(dtsXAudioConfiguration model.DtsXAudioConfiguration) (*model.DtsXAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.DtsXAudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/dtsx", &dtsXAudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete DTS:X Codec Configuration
func (api *EncodingConfigurationsAudioDtsxAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/dtsx/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get DTS:X Codec Configuration Details
func (api *EncodingConfigurationsAudioDtsxAPI) Get(configurationId string) (*model.DtsXAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.DtsXAudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/dtsx/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List DTS:X Codec Configurations
func (api *EncodingConfigurationsAudioDtsxAPI) List(queryParams ...func(*EncodingConfigurationsAudioDtsxAPIListQueryParams)) (*pagination.DtsXAudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioDtsxAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DtsXAudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/dtsx", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioDtsxAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioDtsxAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioDtsxAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
