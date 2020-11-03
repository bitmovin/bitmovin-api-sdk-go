package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioOpusAPI communicates with '/encoding/configurations/audio/opus' endpoints
type EncodingConfigurationsAudioOpusAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/configurations/audio/opus/{configuration_id}/customData' endpoints
    Customdata *EncodingConfigurationsAudioOpusCustomdataAPI

}

// NewEncodingConfigurationsAudioOpusAPI constructor for EncodingConfigurationsAudioOpusAPI that takes options as argument
func NewEncodingConfigurationsAudioOpusAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioOpusAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioOpusAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioOpusAPIWithClient constructor for EncodingConfigurationsAudioOpusAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioOpusAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioOpusAPI {
    a := &EncodingConfigurationsAudioOpusAPI{apiClient: apiClient}
    a.Customdata = NewEncodingConfigurationsAudioOpusCustomdataAPIWithClient(apiClient)

    return a
}

// Create Opus Codec Configuration
func (api *EncodingConfigurationsAudioOpusAPI) Create(opusAudioConfiguration model.OpusAudioConfiguration) (*model.OpusAudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.OpusAudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/opus", &opusAudioConfiguration, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Opus Codec Configuration
func (api *EncodingConfigurationsAudioOpusAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/opus/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Opus Codec Configuration Details
func (api *EncodingConfigurationsAudioOpusAPI) Get(configurationId string) (*model.OpusAudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.OpusAudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/opus/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Opus Configurations
func (api *EncodingConfigurationsAudioOpusAPI) List(queryParams ...func(*EncodingConfigurationsAudioOpusAPIListQueryParams)) (*pagination.OpusAudioConfigurationsListPagination, error) {
    queryParameters := &EncodingConfigurationsAudioOpusAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.OpusAudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/opus", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingConfigurationsAudioOpusAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioOpusAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioOpusAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


