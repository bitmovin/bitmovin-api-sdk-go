package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsMachineLearningAPI intermediary API object with no endpoints
type EncodingEncodingsMachineLearningAPI struct {
    apiClient *apiclient.APIClient

    // ObjectDetection communicates with '/encoding/encodings/{encoding_id}/machine-learning/object-detection' endpoints
    ObjectDetection *EncodingEncodingsMachineLearningObjectDetectionAPI

}

// NewEncodingEncodingsMachineLearningAPI constructor for EncodingEncodingsMachineLearningAPI that takes options as argument
func NewEncodingEncodingsMachineLearningAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMachineLearningAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMachineLearningAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMachineLearningAPIWithClient constructor for EncodingEncodingsMachineLearningAPI that takes an APIClient as argument
func NewEncodingEncodingsMachineLearningAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMachineLearningAPI {
    a := &EncodingEncodingsMachineLearningAPI{apiClient: apiClient}
    a.ObjectDetection = NewEncodingEncodingsMachineLearningObjectDetectionAPIWithClient(apiClient)

    return a
}


