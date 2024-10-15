package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsCaptionsCeaSrtAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/srt' endpoints
type EncodingEncodingsStreamsCaptionsCeaSrtAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/srt/{captions_id}/customData' endpoints
	Customdata *EncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI
}

// NewEncodingEncodingsStreamsCaptionsCeaSrtAPI constructor for EncodingEncodingsStreamsCaptionsCeaSrtAPI that takes options as argument
func NewEncodingEncodingsStreamsCaptionsCeaSrtAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsCaptionsCeaSrtAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsCaptionsCeaSrtAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsCaptionsCeaSrtAPIWithClient constructor for EncodingEncodingsStreamsCaptionsCeaSrtAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsCaptionsCeaSrtAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsCaptionsCeaSrtAPI {
	a := &EncodingEncodingsStreamsCaptionsCeaSrtAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPIWithClient(apiClient)

	return a
}

// Create Embed SRT captions as 608/708 into Stream
func (api *EncodingEncodingsStreamsCaptionsCeaSrtAPI) Create(encodingId string, streamId string, srtToCea608708Caption model.SrtToCea608708Caption) (*model.SrtToCea608708Caption, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.SrtToCea608708Caption
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/srt", &srtToCea608708Caption, &responseModel, reqParams)
	return &responseModel, err
}

// Delete SRT captions as 608/708 from Stream
func (api *EncodingEncodingsStreamsCaptionsCeaSrtAPI) Delete(encodingId string, streamId string, captionsId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["captions_id"] = captionsId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/srt/{captions_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Embed SRT captions as 608/708 Details
func (api *EncodingEncodingsStreamsCaptionsCeaSrtAPI) Get(encodingId string, streamId string, captionsId string) (*model.SrtToCea608708Caption, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["captions_id"] = captionsId
	}

	var responseModel model.SrtToCea608708Caption
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/srt/{captions_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List SRT captions as 608/708 from Stream
func (api *EncodingEncodingsStreamsCaptionsCeaSrtAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsCaptionsCeaSrtAPIListQueryParams)) (*pagination.SrtToCea608708CaptionsListPagination, error) {
	queryParameters := &EncodingEncodingsStreamsCaptionsCeaSrtAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SrtToCea608708CaptionsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/srt", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsStreamsCaptionsCeaSrtAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsCaptionsCeaSrtAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsCaptionsCeaSrtAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
