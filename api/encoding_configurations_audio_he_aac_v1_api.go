package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioHeAacV1API communicates with '/encoding/configurations/audio/he-aac-v1' endpoints
type EncodingConfigurationsAudioHeAacV1API struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/configurations/audio/he-aac-v1/{configuration_id}/customData' endpoints
    Customdata *EncodingConfigurationsAudioHeAacV1CustomdataAPI

}

// NewEncodingConfigurationsAudioHeAacV1API constructor for EncodingConfigurationsAudioHeAacV1API that takes options as argument
func NewEncodingConfigurationsAudioHeAacV1API(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioHeAacV1API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioHeAacV1APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioHeAacV1APIWithClient constructor for EncodingConfigurationsAudioHeAacV1API that takes an APIClient as argument
func NewEncodingConfigurationsAudioHeAacV1APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioHeAacV1API {
    a := &EncodingConfigurationsAudioHeAacV1API{apiClient: apiClient}
    a.Customdata = NewEncodingConfigurationsAudioHeAacV1CustomdataAPIWithClient(apiClient)

    return a
}

// Create HE-AAC v1 Codec Configuration
func (api *EncodingConfigurationsAudioHeAacV1API) Create(heAacV1AudioConfiguration model.HeAacV1AudioConfiguration) (*model.HeAacV1AudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.HeAacV1AudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/he-aac-v1", &heAacV1AudioConfiguration, &responseModel, reqParams)
    return &responseModel, err
}
// Delete HE-AAC v1 Codec Configuration
func (api *EncodingConfigurationsAudioHeAacV1API) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/he-aac-v1/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get HE-AAC v1 Codec Configuration Details
func (api *EncodingConfigurationsAudioHeAacV1API) Get(configurationId string) (*model.HeAacV1AudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.HeAacV1AudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v1/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List HE-AAC v1 Configurations
func (api *EncodingConfigurationsAudioHeAacV1API) List(queryParams ...func(*EncodingConfigurationsAudioHeAacV1APIListQueryParams)) (*pagination.HeAacV1AudioConfigurationsListPagination, error) {
    queryParameters := &EncodingConfigurationsAudioHeAacV1APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.HeAacV1AudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v1", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingConfigurationsAudioHeAacV1APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioHeAacV1APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioHeAacV1APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


