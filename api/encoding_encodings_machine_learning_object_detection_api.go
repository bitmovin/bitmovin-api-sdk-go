package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMachineLearningObjectDetectionAPI communicates with '/encoding/encodings/{encoding_id}/machine-learning/object-detection' endpoints
type EncodingEncodingsMachineLearningObjectDetectionAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/customData' endpoints
	Customdata *EncodingEncodingsMachineLearningObjectDetectionCustomdataAPI
	// Results communicates with '/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/results' endpoints
	Results *EncodingEncodingsMachineLearningObjectDetectionResultsAPI
}

// NewEncodingEncodingsMachineLearningObjectDetectionAPI constructor for EncodingEncodingsMachineLearningObjectDetectionAPI that takes options as argument
func NewEncodingEncodingsMachineLearningObjectDetectionAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMachineLearningObjectDetectionAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMachineLearningObjectDetectionAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMachineLearningObjectDetectionAPIWithClient constructor for EncodingEncodingsMachineLearningObjectDetectionAPI that takes an APIClient as argument
func NewEncodingEncodingsMachineLearningObjectDetectionAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMachineLearningObjectDetectionAPI {
	a := &EncodingEncodingsMachineLearningObjectDetectionAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMachineLearningObjectDetectionCustomdataAPIWithClient(apiClient)
	a.Results = NewEncodingEncodingsMachineLearningObjectDetectionResultsAPIWithClient(apiClient)

	return a
}

// Create Add object detection configuration to an encoding
func (api *EncodingEncodingsMachineLearningObjectDetectionAPI) Create(encodingId string, objectDetectionConfiguration model.ObjectDetectionConfiguration) (*model.ObjectDetectionConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.ObjectDetectionConfiguration
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/machine-learning/object-detection", &objectDetectionConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete object detection configuration
func (api *EncodingEncodingsMachineLearningObjectDetectionAPI) Delete(encodingId string, objectDetectionId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["object_detection_id"] = objectDetectionId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get object detection configuration details
func (api *EncodingEncodingsMachineLearningObjectDetectionAPI) Get(encodingId string, objectDetectionId string) (*model.ObjectDetectionConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["object_detection_id"] = objectDetectionId
	}

	var responseModel model.ObjectDetectionConfiguration
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List object detection configurations of an encoding
func (api *EncodingEncodingsMachineLearningObjectDetectionAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMachineLearningObjectDetectionAPIListQueryParams)) (*pagination.ObjectDetectionConfigurationsListPagination, error) {
	queryParameters := &EncodingEncodingsMachineLearningObjectDetectionAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ObjectDetectionConfigurationsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMachineLearningObjectDetectionAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMachineLearningObjectDetectionAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMachineLearningObjectDetectionAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
