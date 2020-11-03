package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsGenericS3CustomdataAPI communicates with '/encoding/outputs/generic-s3/{output_id}/customData' endpoints
type EncodingOutputsGenericS3CustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingOutputsGenericS3CustomdataAPI constructor for EncodingOutputsGenericS3CustomdataAPI that takes options as argument
func NewEncodingOutputsGenericS3CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsGenericS3CustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsGenericS3CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsGenericS3CustomdataAPIWithClient constructor for EncodingOutputsGenericS3CustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsGenericS3CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsGenericS3CustomdataAPI {
    a := &EncodingOutputsGenericS3CustomdataAPI{apiClient: apiClient}
    return a
}

// Get Generic S3 Output Custom Data
func (api *EncodingOutputsGenericS3CustomdataAPI) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/outputs/generic-s3/{output_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

