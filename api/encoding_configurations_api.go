package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAPI communicates with '/encoding/configurations' endpoints
type EncodingConfigurationsAPI struct {
	apiClient *apiclient.APIClient

	// Type communicates with '/encoding/configurations/{configuration_id}/type' endpoints
	Type *EncodingConfigurationsTypeAPI
	// Video intermediary API object with no endpoints
	Video *EncodingConfigurationsVideoAPI
	// Audio intermediary API object with no endpoints
	Audio *EncodingConfigurationsAudioAPI
	// Subtitles intermediary API object with no endpoints
	Subtitles *EncodingConfigurationsSubtitlesAPI
}

// NewEncodingConfigurationsAPI constructor for EncodingConfigurationsAPI that takes options as argument
func NewEncodingConfigurationsAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAPIWithClient constructor for EncodingConfigurationsAPI that takes an APIClient as argument
func NewEncodingConfigurationsAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAPI {
	a := &EncodingConfigurationsAPI{apiClient: apiClient}
	a.Type = NewEncodingConfigurationsTypeAPIWithClient(apiClient)
	a.Video = NewEncodingConfigurationsVideoAPIWithClient(apiClient)
	a.Audio = NewEncodingConfigurationsAudioAPIWithClient(apiClient)
	a.Subtitles = NewEncodingConfigurationsSubtitlesAPIWithClient(apiClient)

	return a
}

// Get Codec Configuration Details
func (api *EncodingConfigurationsAPI) Get(configurationId string) (*model.CodecConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CodecConfiguration
	err := api.apiClient.Get("/encoding/configurations/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Codec Configurations
func (api *EncodingConfigurationsAPI) List(queryParams ...func(*EncodingConfigurationsAPIListQueryParams)) (*pagination.CodecConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.CodecConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
