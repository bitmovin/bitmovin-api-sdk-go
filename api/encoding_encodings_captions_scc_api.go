package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsCaptionsSccAPI communicates with '/encoding/encodings/{encoding_id}/captions/scc' endpoints
type EncodingEncodingsCaptionsSccAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/captions/scc/{captions_id}/customData' endpoints
    Customdata *EncodingEncodingsCaptionsSccCustomdataAPI

}

// NewEncodingEncodingsCaptionsSccAPI constructor for EncodingEncodingsCaptionsSccAPI that takes options as argument
func NewEncodingEncodingsCaptionsSccAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsCaptionsSccAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsCaptionsSccAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsCaptionsSccAPIWithClient constructor for EncodingEncodingsCaptionsSccAPI that takes an APIClient as argument
func NewEncodingEncodingsCaptionsSccAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsCaptionsSccAPI {
    a := &EncodingEncodingsCaptionsSccAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsCaptionsSccCustomdataAPIWithClient(apiClient)

    return a
}

// Create Convert SCC captions
func (api *EncodingEncodingsCaptionsSccAPI) Create(encodingId string, convertSccCaption model.ConvertSccCaption) (*model.ConvertSccCaption, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.ConvertSccCaption
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/captions/scc", &convertSccCaption, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Convert SCC captions
func (api *EncodingEncodingsCaptionsSccAPI) Delete(encodingId string, captionsId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/captions/scc/{captions_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Convert SCC captions Details
func (api *EncodingEncodingsCaptionsSccAPI) Get(encodingId string, captionsId string) (*model.ConvertSccCaption, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
    }

    var responseModel model.ConvertSccCaption
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/scc/{captions_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Convert SCC captions
func (api *EncodingEncodingsCaptionsSccAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsCaptionsSccAPIListQueryParams)) (*pagination.ConvertSccCaptionsListPagination, error) {
    queryParameters := &EncodingEncodingsCaptionsSccAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ConvertSccCaptionsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/scc", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsCaptionsSccAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsCaptionsSccAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsCaptionsSccAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


