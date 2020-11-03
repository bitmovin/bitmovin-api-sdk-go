package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsS3CustomdataAPI communicates with '/encoding/inputs/s3/{input_id}/customData' endpoints
type EncodingInputsS3CustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsS3CustomdataAPI constructor for EncodingInputsS3CustomdataAPI that takes options as argument
func NewEncodingInputsS3CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsS3CustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsS3CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsS3CustomdataAPIWithClient constructor for EncodingInputsS3CustomdataAPI that takes an APIClient as argument
func NewEncodingInputsS3CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsS3CustomdataAPI {
    a := &EncodingInputsS3CustomdataAPI{apiClient: apiClient}
    return a
}

// Get S3 Custom Data
func (api *EncodingInputsS3CustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/s3/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

