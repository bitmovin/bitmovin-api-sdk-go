package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMachineLearningObjectDetectionResultsAPI communicates with '/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/results' endpoints
type EncodingEncodingsMachineLearningObjectDetectionResultsAPI struct {
	apiClient *apiclient.APIClient

	// ByTimestamp communicates with '/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/results/by-timestamp' endpoints
	ByTimestamp *EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPI
}

// NewEncodingEncodingsMachineLearningObjectDetectionResultsAPI constructor for EncodingEncodingsMachineLearningObjectDetectionResultsAPI that takes options as argument
func NewEncodingEncodingsMachineLearningObjectDetectionResultsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMachineLearningObjectDetectionResultsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMachineLearningObjectDetectionResultsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMachineLearningObjectDetectionResultsAPIWithClient constructor for EncodingEncodingsMachineLearningObjectDetectionResultsAPI that takes an APIClient as argument
func NewEncodingEncodingsMachineLearningObjectDetectionResultsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMachineLearningObjectDetectionResultsAPI {
	a := &EncodingEncodingsMachineLearningObjectDetectionResultsAPI{apiClient: apiClient}
	a.ByTimestamp = NewEncodingEncodingsMachineLearningObjectDetectionResultsByTimestampAPIWithClient(apiClient)

	return a
}

// List object detection results
func (api *EncodingEncodingsMachineLearningObjectDetectionResultsAPI) List(encodingId string, objectDetectionId string, queryParams ...func(*EncodingEncodingsMachineLearningObjectDetectionResultsAPIListQueryParams)) (*pagination.ObjectDetectionResultsListPagination, error) {
	queryParameters := &EncodingEncodingsMachineLearningObjectDetectionResultsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["object_detection_id"] = objectDetectionId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ObjectDetectionResultsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/results", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMachineLearningObjectDetectionResultsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMachineLearningObjectDetectionResultsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMachineLearningObjectDetectionResultsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
