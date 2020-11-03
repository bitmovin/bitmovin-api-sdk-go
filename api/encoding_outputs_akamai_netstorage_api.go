package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsAkamaiNetstorageAPI communicates with '/encoding/outputs/akamai-netstorage' endpoints
type EncodingOutputsAkamaiNetstorageAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/outputs/akamai-netstorage/{output_id}/customData' endpoints
    Customdata *EncodingOutputsAkamaiNetstorageCustomdataAPI

}

// NewEncodingOutputsAkamaiNetstorageAPI constructor for EncodingOutputsAkamaiNetstorageAPI that takes options as argument
func NewEncodingOutputsAkamaiNetstorageAPI(options ...apiclient.APIClientOption) (*EncodingOutputsAkamaiNetstorageAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsAkamaiNetstorageAPIWithClient(apiClient), nil
}

// NewEncodingOutputsAkamaiNetstorageAPIWithClient constructor for EncodingOutputsAkamaiNetstorageAPI that takes an APIClient as argument
func NewEncodingOutputsAkamaiNetstorageAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsAkamaiNetstorageAPI {
    a := &EncodingOutputsAkamaiNetstorageAPI{apiClient: apiClient}
    a.Customdata = NewEncodingOutputsAkamaiNetstorageCustomdataAPIWithClient(apiClient)

    return a
}

// Create Akamai NetStorage Output
func (api *EncodingOutputsAkamaiNetstorageAPI) Create(akamaiNetStorageOutput model.AkamaiNetStorageOutput) (*model.AkamaiNetStorageOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AkamaiNetStorageOutput
    err := api.apiClient.Post("/encoding/outputs/akamai-netstorage", &akamaiNetStorageOutput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Akamai NetStorage Output
func (api *EncodingOutputsAkamaiNetstorageAPI) Delete(outputId string) (*model.AkamaiNetStorageOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.AkamaiNetStorageOutput
    err := api.apiClient.Delete("/encoding/outputs/akamai-netstorage/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Akamai NetStorage Output Details
func (api *EncodingOutputsAkamaiNetstorageAPI) Get(outputId string) (*model.AkamaiNetStorageOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.AkamaiNetStorageOutput
    err := api.apiClient.Get("/encoding/outputs/akamai-netstorage/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Akamai NetStorage Outputs
func (api *EncodingOutputsAkamaiNetstorageAPI) List(queryParams ...func(*EncodingOutputsAkamaiNetstorageAPIListQueryParams)) (*pagination.AkamaiNetStorageOutputsListPagination, error) {
    queryParameters := &EncodingOutputsAkamaiNetstorageAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.AkamaiNetStorageOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/akamai-netstorage", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingOutputsAkamaiNetstorageAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsAkamaiNetstorageAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsAkamaiNetstorageAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


