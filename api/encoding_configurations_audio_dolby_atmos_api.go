package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioDolbyAtmosAPI communicates with '/encoding/configurations/audio/dolby-atmos' endpoints
type EncodingConfigurationsAudioDolbyAtmosAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/configurations/audio/dolby-atmos/{configuration_id}/customData' endpoints
    Customdata *EncodingConfigurationsAudioDolbyAtmosCustomdataAPI

}

// NewEncodingConfigurationsAudioDolbyAtmosAPI constructor for EncodingConfigurationsAudioDolbyAtmosAPI that takes options as argument
func NewEncodingConfigurationsAudioDolbyAtmosAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDolbyAtmosAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioDolbyAtmosAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDolbyAtmosAPIWithClient constructor for EncodingConfigurationsAudioDolbyAtmosAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDolbyAtmosAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDolbyAtmosAPI {
    a := &EncodingConfigurationsAudioDolbyAtmosAPI{apiClient: apiClient}
    a.Customdata = NewEncodingConfigurationsAudioDolbyAtmosCustomdataAPIWithClient(apiClient)

    return a
}

// Create Dolby Atmos Configuration
func (api *EncodingConfigurationsAudioDolbyAtmosAPI) Create(dolbyAtmosAudioConfiguration model.DolbyAtmosAudioConfiguration) (*model.DolbyAtmosAudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.DolbyAtmosAudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/dolby-atmos", &dolbyAtmosAudioConfiguration, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Dolby Atmos Codec Configuration
func (api *EncodingConfigurationsAudioDolbyAtmosAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/dolby-atmos/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Dolby Atmos Codec Configuration Details
func (api *EncodingConfigurationsAudioDolbyAtmosAPI) Get(configurationId string) (*model.DolbyAtmosAudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.DolbyAtmosAudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/dolby-atmos/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Dolby Atmos Configurations
func (api *EncodingConfigurationsAudioDolbyAtmosAPI) List(queryParams ...func(*EncodingConfigurationsAudioDolbyAtmosAPIListQueryParams)) (*pagination.DolbyAtmosAudioConfigurationsListPagination, error) {
    queryParameters := &EncodingConfigurationsAudioDolbyAtmosAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DolbyAtmosAudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/dolby-atmos", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingConfigurationsAudioDolbyAtmosAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioDolbyAtmosAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioDolbyAtmosAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


