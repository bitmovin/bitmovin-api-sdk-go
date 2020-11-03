package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsSubtitlesWebvttAPI communicates with '/encoding/configurations/subtitles/webvtt/' endpoints
type EncodingConfigurationsSubtitlesWebvttAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/configurations/subtitles/webvtt/{configuration_id}/customData' endpoints
    Customdata *EncodingConfigurationsSubtitlesWebvttCustomdataAPI

}

// NewEncodingConfigurationsSubtitlesWebvttAPI constructor for EncodingConfigurationsSubtitlesWebvttAPI that takes options as argument
func NewEncodingConfigurationsSubtitlesWebvttAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsSubtitlesWebvttAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsSubtitlesWebvttAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsSubtitlesWebvttAPIWithClient constructor for EncodingConfigurationsSubtitlesWebvttAPI that takes an APIClient as argument
func NewEncodingConfigurationsSubtitlesWebvttAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsSubtitlesWebvttAPI {
    a := &EncodingConfigurationsSubtitlesWebvttAPI{apiClient: apiClient}
    a.Customdata = NewEncodingConfigurationsSubtitlesWebvttCustomdataAPIWithClient(apiClient)

    return a
}

// Create WebVtt Subtitle Configuration
func (api *EncodingConfigurationsSubtitlesWebvttAPI) Create(webVttConfiguration model.WebVttConfiguration) (*model.WebVttConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.WebVttConfiguration
    err := api.apiClient.Post("/encoding/configurations/subtitles/webvtt/", &webVttConfiguration, &responseModel, reqParams)
    return &responseModel, err
}
// Delete WebVtt Subtitle Configuration
func (api *EncodingConfigurationsSubtitlesWebvttAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/subtitles/webvtt/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get WebVtt Subtitle Configuration Details
func (api *EncodingConfigurationsSubtitlesWebvttAPI) Get(configurationId string) (*model.WebVttConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.WebVttConfiguration
    err := api.apiClient.Get("/encoding/configurations/subtitles/webvtt/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List WebVtt Configurations
func (api *EncodingConfigurationsSubtitlesWebvttAPI) List(queryParams ...func(*EncodingConfigurationsSubtitlesWebvttAPIListQueryParams)) (*pagination.WebVttConfigurationsListPagination, error) {
    queryParameters := &EncodingConfigurationsSubtitlesWebvttAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.WebVttConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/subtitles/webvtt/", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingConfigurationsSubtitlesWebvttAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsSubtitlesWebvttAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsSubtitlesWebvttAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


