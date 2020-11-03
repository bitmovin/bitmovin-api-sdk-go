package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsVideoVp9API communicates with '/encoding/configurations/video/vp9' endpoints
type EncodingConfigurationsVideoVp9API struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/configurations/video/vp9/{configuration_id}/customData' endpoints
    Customdata *EncodingConfigurationsVideoVp9CustomdataAPI

}

// NewEncodingConfigurationsVideoVp9API constructor for EncodingConfigurationsVideoVp9API that takes options as argument
func NewEncodingConfigurationsVideoVp9API(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoVp9API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsVideoVp9APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoVp9APIWithClient constructor for EncodingConfigurationsVideoVp9API that takes an APIClient as argument
func NewEncodingConfigurationsVideoVp9APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoVp9API {
    a := &EncodingConfigurationsVideoVp9API{apiClient: apiClient}
    a.Customdata = NewEncodingConfigurationsVideoVp9CustomdataAPIWithClient(apiClient)

    return a
}

// Create VP9 Codec Configuration
func (api *EncodingConfigurationsVideoVp9API) Create(vp9VideoConfiguration model.Vp9VideoConfiguration) (*model.Vp9VideoConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.Vp9VideoConfiguration
    err := api.apiClient.Post("/encoding/configurations/video/vp9", &vp9VideoConfiguration, &responseModel, reqParams)
    return &responseModel, err
}
// Delete VP9 Codec Configuration
func (api *EncodingConfigurationsVideoVp9API) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/video/vp9/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get VP9 Codec Configuration Details
func (api *EncodingConfigurationsVideoVp9API) Get(configurationId string) (*model.Vp9VideoConfiguration, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.Vp9VideoConfiguration
    err := api.apiClient.Get("/encoding/configurations/video/vp9/{configuration_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List VP9 Codec Configurations
func (api *EncodingConfigurationsVideoVp9API) List(queryParams ...func(*EncodingConfigurationsVideoVp9APIListQueryParams)) (*pagination.Vp9VideoConfigurationsListPagination, error) {
    queryParameters := &EncodingConfigurationsVideoVp9APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.Vp9VideoConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/video/vp9", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingConfigurationsVideoVp9APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsVideoVp9APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsVideoVp9APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


