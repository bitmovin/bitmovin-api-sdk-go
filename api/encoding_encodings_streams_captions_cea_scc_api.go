package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsCaptionsCeaSccAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc' endpoints
type EncodingEncodingsStreamsCaptionsCeaSccAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc/{captions_id}/customData' endpoints
    Customdata *EncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI

}

// NewEncodingEncodingsStreamsCaptionsCeaSccAPI constructor for EncodingEncodingsStreamsCaptionsCeaSccAPI that takes options as argument
func NewEncodingEncodingsStreamsCaptionsCeaSccAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsCaptionsCeaSccAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsCaptionsCeaSccAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsCaptionsCeaSccAPIWithClient constructor for EncodingEncodingsStreamsCaptionsCeaSccAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsCaptionsCeaSccAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsCaptionsCeaSccAPI {
    a := &EncodingEncodingsStreamsCaptionsCeaSccAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsStreamsCaptionsCeaSccCustomdataAPIWithClient(apiClient)

    return a
}

// Create Embed SCC captions as 608/708 into Stream
func (api *EncodingEncodingsStreamsCaptionsCeaSccAPI) Create(encodingId string, streamId string, sccCaption model.SccCaption) (*model.SccCaption, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel model.SccCaption
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc", &sccCaption, &responseModel, reqParams)
    return &responseModel, err
}
// Delete SCC captions as 608/708 from Stream
func (api *EncodingEncodingsStreamsCaptionsCeaSccAPI) Delete(encodingId string, streamId string, captionsId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["captions_id"] = captionsId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc/{captions_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Embed SCC captions as 608/708 Details
func (api *EncodingEncodingsStreamsCaptionsCeaSccAPI) Get(encodingId string, streamId string, captionsId string) (*model.SccCaption, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["captions_id"] = captionsId
    }

    var responseModel model.SccCaption
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc/{captions_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List SCC captions as 608/708 from Stream
func (api *EncodingEncodingsStreamsCaptionsCeaSccAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsCaptionsCeaSccAPIListQueryParams)) (*pagination.SccCaptionsListPagination, error) {
    queryParameters := &EncodingEncodingsStreamsCaptionsCeaSccAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SccCaptionsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsStreamsCaptionsCeaSccAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsCaptionsCeaSccAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsCaptionsCeaSccAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


