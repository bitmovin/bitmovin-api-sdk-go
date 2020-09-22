package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest' endpoints
type EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI constructor for EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI that takes options as argument
func NewEncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPIWithClient constructor for EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI {
	a := &EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI{apiClient: apiClient}
	return a
}

// Create Add Dolby Vision Metadata Ingest Input Stream
func (api *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI) Create(encodingId string, dolbyVisionMetadataIngestInputStream model.DolbyVisionMetadataIngestInputStream) (*model.DolbyVisionMetadataIngestInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.DolbyVisionMetadataIngestInputStream
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest", &dolbyVisionMetadataIngestInputStream, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Dolby Vision Metadata Ingest Input Stream
func (api *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Dolby Vision Metadata Ingest Input Stream Details
func (api *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI) Get(encodingId string, inputStreamId string) (*model.DolbyVisionMetadataIngestInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.DolbyVisionMetadataIngestInputStream
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Dolby Vision Metadata Ingest Input Streams
func (api *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPIListQueryParams)) (*pagination.DolbyVisionMetadataIngestInputStreamsListPagination, error) {
	queryParameters := &EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DolbyVisionMetadataIngestInputStreamsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
