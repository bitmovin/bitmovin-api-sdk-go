package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMachineLearningObjectDetectionCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/customData' endpoints
type EncodingEncodingsMachineLearningObjectDetectionCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMachineLearningObjectDetectionCustomdataAPI constructor for EncodingEncodingsMachineLearningObjectDetectionCustomdataAPI that takes options as argument
func NewEncodingEncodingsMachineLearningObjectDetectionCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMachineLearningObjectDetectionCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMachineLearningObjectDetectionCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMachineLearningObjectDetectionCustomdataAPIWithClient constructor for EncodingEncodingsMachineLearningObjectDetectionCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMachineLearningObjectDetectionCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMachineLearningObjectDetectionCustomdataAPI {
    a := &EncodingEncodingsMachineLearningObjectDetectionCustomdataAPI{apiClient: apiClient}
    return a
}

// Get the custom data of an object detection configuration
func (api *EncodingEncodingsMachineLearningObjectDetectionCustomdataAPI) Get(encodingId string, objectDetectionId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["object_detection_id"] = objectDetectionId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

