package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsS3CustomdataAPI communicates with '/encoding/outputs/s3/{output_id}/customData' endpoints
type EncodingOutputsS3CustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingOutputsS3CustomdataAPI constructor for EncodingOutputsS3CustomdataAPI that takes options as argument
func NewEncodingOutputsS3CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsS3CustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsS3CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsS3CustomdataAPIWithClient constructor for EncodingOutputsS3CustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsS3CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsS3CustomdataAPI {
    a := &EncodingOutputsS3CustomdataAPI{apiClient: apiClient}
    return a
}

// Get S3 Output Custom Data
func (api *EncodingOutputsS3CustomdataAPI) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/outputs/s3/{output_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

