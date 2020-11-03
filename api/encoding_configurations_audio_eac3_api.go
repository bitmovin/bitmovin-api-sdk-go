package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioEac3API communicates with '/encoding/configurations/audio/eac3' endpoints
type EncodingConfigurationsAudioEac3API struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/configurations/audio/eac3/{configuration_id}/customData' endpoints
    Customdata *EncodingConfigurationsAudioEac3CustomdataAPI

}

// NewEncodingConfigurationsAudioEac3API constructor for EncodingConfigurationsAudioEac3API that takes options as argument
func NewEncodingConfigurationsAudioEac3API(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioEac3API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioEac3APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioEac3APIWithClient constructor for EncodingConfigurationsAudioEac3API that takes an APIClient as argument
func NewEncodingConfigurationsAudioEac3APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioEac3API {
    a := &EncodingConfigurationsAudioEac3API{apiClient: apiClient}
    a.Customdata = NewEncodingConfigurationsAudioEac3CustomdataAPIWithClient(apiClient)

    return a
}

// Create E-AC3 Codec Configuration
func (api *EncodingConfigurationsAudioEac3API) Create(eac3AudioConfiguration model.Eac3AudioConfiguration) (*model.Eac3AudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.Eac3AudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/eac3", &eac3AudioConfiguration, &responseModel, reqParams)
    return &responseModel, err
}
// Delete E-AC3 Codec Configuration
func (api *EncodingConfigurationsAudioEac3API) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/eac3/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get E-AC3 Codec Configuration Details
func (api *EncodingConfigurationsAudioEac3API) Get(configurationId string) (*model.Eac3AudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.Eac3AudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/eac3/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List E-AC3 Configurations
func (api *EncodingConfigurationsAudioEac3API) List(queryParams ...func(*EncodingConfigurationsAudioEac3APIListQueryParams)) (*pagination.Eac3AudioConfigurationsListPagination, error) {
    queryParameters := &EncodingConfigurationsAudioEac3APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.Eac3AudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/eac3", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingConfigurationsAudioEac3APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioEac3APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioEac3APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


