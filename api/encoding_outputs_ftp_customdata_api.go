package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsFtpCustomdataAPI communicates with '/encoding/outputs/ftp/{output_id}/customData' endpoints
type EncodingOutputsFtpCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingOutputsFtpCustomdataAPI constructor for EncodingOutputsFtpCustomdataAPI that takes options as argument
func NewEncodingOutputsFtpCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsFtpCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsFtpCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsFtpCustomdataAPIWithClient constructor for EncodingOutputsFtpCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsFtpCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsFtpCustomdataAPI {
    a := &EncodingOutputsFtpCustomdataAPI{apiClient: apiClient}
    return a
}

// Get FTP Output Custom Data
func (api *EncodingOutputsFtpCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/outputs/ftp/{output_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

