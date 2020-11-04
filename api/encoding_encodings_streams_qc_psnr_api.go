package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsQcPsnrAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/qc/psnr' endpoints
type EncodingEncodingsStreamsQcPsnrAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsQcPsnrAPI constructor for EncodingEncodingsStreamsQcPsnrAPI that takes options as argument
func NewEncodingEncodingsStreamsQcPsnrAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsQcPsnrAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsQcPsnrAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsQcPsnrAPIWithClient constructor for EncodingEncodingsStreamsQcPsnrAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsQcPsnrAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsQcPsnrAPI {
	a := &EncodingEncodingsStreamsQcPsnrAPI{apiClient: apiClient}
	return a
}

// Create Activate PSNR quality metrics for the selected stream
func (api *EncodingEncodingsStreamsQcPsnrAPI) Create(encodingId string, streamId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/qc/psnr", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Get Stream PSNR metrics
func (api *EncodingEncodingsStreamsQcPsnrAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsQcPsnrAPIListQueryParams)) (*pagination.PsnrQualityMetricsListPagination, error) {
	queryParameters := &EncodingEncodingsStreamsQcPsnrAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.PsnrQualityMetricsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/qc/psnr", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsStreamsQcPsnrAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsQcPsnrAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsQcPsnrAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
