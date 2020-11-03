package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsLiveMediaIngestAPI communicates with '/encoding/outputs/live-media-ingest' endpoints
type EncodingOutputsLiveMediaIngestAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/outputs/live-media-ingest/{output_id}/customData' endpoints
    Customdata *EncodingOutputsLiveMediaIngestCustomdataAPI

}

// NewEncodingOutputsLiveMediaIngestAPI constructor for EncodingOutputsLiveMediaIngestAPI that takes options as argument
func NewEncodingOutputsLiveMediaIngestAPI(options ...apiclient.APIClientOption) (*EncodingOutputsLiveMediaIngestAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsLiveMediaIngestAPIWithClient(apiClient), nil
}

// NewEncodingOutputsLiveMediaIngestAPIWithClient constructor for EncodingOutputsLiveMediaIngestAPI that takes an APIClient as argument
func NewEncodingOutputsLiveMediaIngestAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsLiveMediaIngestAPI {
    a := &EncodingOutputsLiveMediaIngestAPI{apiClient: apiClient}
    a.Customdata = NewEncodingOutputsLiveMediaIngestCustomdataAPIWithClient(apiClient)

    return a
}

// Create Live Media Ingest Output
func (api *EncodingOutputsLiveMediaIngestAPI) Create(liveMediaIngestOutput model.LiveMediaIngestOutput) (*model.LiveMediaIngestOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.LiveMediaIngestOutput
    err := api.apiClient.Post("/encoding/outputs/live-media-ingest", &liveMediaIngestOutput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Live Media Ingest Output
func (api *EncodingOutputsLiveMediaIngestAPI) Delete(outputId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/outputs/live-media-ingest/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Live Media Ingest Output Details
func (api *EncodingOutputsLiveMediaIngestAPI) Get(outputId string) (*model.LiveMediaIngestOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.LiveMediaIngestOutput
    err := api.apiClient.Get("/encoding/outputs/live-media-ingest/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Live Media Ingest Outputs
func (api *EncodingOutputsLiveMediaIngestAPI) List(queryParams ...func(*EncodingOutputsLiveMediaIngestAPIListQueryParams)) (*pagination.LiveMediaIngestOutputsListPagination, error) {
    queryParameters := &EncodingOutputsLiveMediaIngestAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.LiveMediaIngestOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/live-media-ingest", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingOutputsLiveMediaIngestAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsLiveMediaIngestAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsLiveMediaIngestAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


