package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioHeAacV2API communicates with '/encoding/configurations/audio/he-aac-v2' endpoints
type EncodingConfigurationsAudioHeAacV2API struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/configurations/audio/he-aac-v2/{configuration_id}/customData' endpoints
    Customdata *EncodingConfigurationsAudioHeAacV2CustomdataAPI

}

// NewEncodingConfigurationsAudioHeAacV2API constructor for EncodingConfigurationsAudioHeAacV2API that takes options as argument
func NewEncodingConfigurationsAudioHeAacV2API(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioHeAacV2API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioHeAacV2APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioHeAacV2APIWithClient constructor for EncodingConfigurationsAudioHeAacV2API that takes an APIClient as argument
func NewEncodingConfigurationsAudioHeAacV2APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioHeAacV2API {
    a := &EncodingConfigurationsAudioHeAacV2API{apiClient: apiClient}
    a.Customdata = NewEncodingConfigurationsAudioHeAacV2CustomdataAPIWithClient(apiClient)

    return a
}

// Create HE-AAC v2 Codec Configuration
func (api *EncodingConfigurationsAudioHeAacV2API) Create(heAacV2AudioConfiguration model.HeAacV2AudioConfiguration) (*model.HeAacV2AudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.HeAacV2AudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/he-aac-v2", &heAacV2AudioConfiguration, &responseModel, reqParams)
    return &responseModel, err
}
// Delete HE-AAC v2 Codec Configuration
func (api *EncodingConfigurationsAudioHeAacV2API) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/he-aac-v2/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get HE-AAC v2 Codec Configuration Details
func (api *EncodingConfigurationsAudioHeAacV2API) Get(configurationId string) (*model.HeAacV2AudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.HeAacV2AudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v2/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List HE-AAC v2 Configurations
func (api *EncodingConfigurationsAudioHeAacV2API) List(queryParams ...func(*EncodingConfigurationsAudioHeAacV2APIListQueryParams)) (*pagination.HeAacV2AudioConfigurationsListPagination, error) {
    queryParameters := &EncodingConfigurationsAudioHeAacV2APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.HeAacV2AudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v2", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingConfigurationsAudioHeAacV2APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioHeAacV2APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioHeAacV2APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


