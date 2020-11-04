package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI communicates with '/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/results/by-timestamp' endpoints
type EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI constructor for EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI that takes options as argument
func NewEncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPIWithClient constructor for EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI that takes an APIClient as argument
func NewEncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI {
	a := &EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI{apiClient: apiClient}
	return a
}

// List object detection results grouped by timestamp
func (api *EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI) List(encodingId string, objectDetectionId string, queryParams ...func(*EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPIListQueryParams)) (*pagination.ObjectDetectionTimestampResultsListPagination, error) {
	queryParameters := &EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["object_detection_id"] = objectDetectionId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ObjectDetectionTimestampResultsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/results/by-timestamp", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
