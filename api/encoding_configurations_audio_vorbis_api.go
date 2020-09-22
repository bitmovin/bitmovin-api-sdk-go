package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioVorbisAPI communicates with '/encoding/configurations/audio/vorbis' endpoints
type EncodingConfigurationsAudioVorbisAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/vorbis/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioVorbisCustomdataAPI
}

// NewEncodingConfigurationsAudioVorbisAPI constructor for EncodingConfigurationsAudioVorbisAPI that takes options as argument
func NewEncodingConfigurationsAudioVorbisAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioVorbisAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioVorbisAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioVorbisAPIWithClient constructor for EncodingConfigurationsAudioVorbisAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioVorbisAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioVorbisAPI {
	a := &EncodingConfigurationsAudioVorbisAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioVorbisCustomdataAPIWithClient(apiClient)

	return a
}

// Create Vorbis Codec Configuration
func (api *EncodingConfigurationsAudioVorbisAPI) Create(vorbisAudioConfiguration model.VorbisAudioConfiguration) (*model.VorbisAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.VorbisAudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/vorbis", &vorbisAudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Vorbis Codec Configuration
func (api *EncodingConfigurationsAudioVorbisAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/vorbis/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Vorbis Codec Configuration Details
func (api *EncodingConfigurationsAudioVorbisAPI) Get(configurationId string) (*model.VorbisAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.VorbisAudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/vorbis/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Vorbis Configurations
func (api *EncodingConfigurationsAudioVorbisAPI) List(queryParams ...func(*EncodingConfigurationsAudioVorbisAPIListQueryParams)) (*pagination.VorbisAudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioVorbisAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.VorbisAudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/vorbis", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioVorbisAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioVorbisAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioVorbisAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
