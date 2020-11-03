package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsVideoVp8API communicates with '/encoding/configurations/video/vp8' endpoints
type EncodingConfigurationsVideoVp8API struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/configurations/video/vp8/{configuration_id}/customData' endpoints
    Customdata *EncodingConfigurationsVideoVp8CustomdataAPI

}

// NewEncodingConfigurationsVideoVp8API constructor for EncodingConfigurationsVideoVp8API that takes options as argument
func NewEncodingConfigurationsVideoVp8API(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoVp8API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsVideoVp8APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoVp8APIWithClient constructor for EncodingConfigurationsVideoVp8API that takes an APIClient as argument
func NewEncodingConfigurationsVideoVp8APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoVp8API {
    a := &EncodingConfigurationsVideoVp8API{apiClient: apiClient}
    a.Customdata = NewEncodingConfigurationsVideoVp8CustomdataAPIWithClient(apiClient)

    return a
}

// Create VP8 Codec Configuration
func (api *EncodingConfigurationsVideoVp8API) Create(vp8VideoConfiguration model.Vp8VideoConfiguration) (*model.Vp8VideoConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.Vp8VideoConfiguration
    err := api.apiClient.Post("/encoding/configurations/video/vp8", &vp8VideoConfiguration, &responseModel, reqParams)
    return &responseModel, err
}
// Delete VP8 Codec Configuration
func (api *EncodingConfigurationsVideoVp8API) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/video/vp8/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get VP8 Codec Configuration Details
func (api *EncodingConfigurationsVideoVp8API) Get(configurationId string) (*model.Vp8VideoConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.Vp8VideoConfiguration
    err := api.apiClient.Get("/encoding/configurations/video/vp8/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List issues a Get request to /encoding/configurations/video/vp8
func (api *EncodingConfigurationsVideoVp8API) List(queryParams ...func(*EncodingConfigurationsVideoVp8APIListQueryParams)) (*pagination.Vp8VideoConfigurationsListPagination, error) {
    queryParameters := &EncodingConfigurationsVideoVp8APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.Vp8VideoConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/video/vp8", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingConfigurationsVideoVp8APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsVideoVp8APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsVideoVp8APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


