package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioDtsPassthroughAPI communicates with '/encoding/configurations/audio/dts-passthrough' endpoints
type EncodingConfigurationsAudioDtsPassthroughAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/configurations/audio/dts-passthrough/{configuration_id}/customData' endpoints
    Customdata *EncodingConfigurationsAudioDtsPassthroughCustomdataAPI

}

// NewEncodingConfigurationsAudioDtsPassthroughAPI constructor for EncodingConfigurationsAudioDtsPassthroughAPI that takes options as argument
func NewEncodingConfigurationsAudioDtsPassthroughAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDtsPassthroughAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioDtsPassthroughAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDtsPassthroughAPIWithClient constructor for EncodingConfigurationsAudioDtsPassthroughAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDtsPassthroughAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDtsPassthroughAPI {
    a := &EncodingConfigurationsAudioDtsPassthroughAPI{apiClient: apiClient}
    a.Customdata = NewEncodingConfigurationsAudioDtsPassthroughCustomdataAPIWithClient(apiClient)

    return a
}

// Create DTS Passthrough Configuration
func (api *EncodingConfigurationsAudioDtsPassthroughAPI) Create(dtsPassthroughAudioConfiguration model.DtsPassthroughAudioConfiguration) (*model.DtsPassthroughAudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.DtsPassthroughAudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/dts-passthrough", &dtsPassthroughAudioConfiguration, &responseModel, reqParams)
    return &responseModel, err
}
// Delete DTS Passthrough Codec Configuration
func (api *EncodingConfigurationsAudioDtsPassthroughAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/dts-passthrough/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get DTS Passthrough Codec Configuration Details
func (api *EncodingConfigurationsAudioDtsPassthroughAPI) Get(configurationId string) (*model.DtsPassthroughAudioConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.DtsPassthroughAudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/dts-passthrough/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List DTS Passthrough Configurations
func (api *EncodingConfigurationsAudioDtsPassthroughAPI) List(queryParams ...func(*EncodingConfigurationsAudioDtsPassthroughAPIListQueryParams)) (*pagination.DtsPassthroughAudioConfigurationsListPagination, error) {
    queryParameters := &EncodingConfigurationsAudioDtsPassthroughAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DtsPassthroughAudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/dts-passthrough", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingConfigurationsAudioDtsPassthroughAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioDtsPassthroughAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioDtsPassthroughAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


