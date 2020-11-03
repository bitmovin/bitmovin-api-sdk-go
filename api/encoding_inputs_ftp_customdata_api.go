package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsFtpCustomdataAPI communicates with '/encoding/inputs/ftp/{input_id}/customData' endpoints
type EncodingInputsFtpCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsFtpCustomdataAPI constructor for EncodingInputsFtpCustomdataAPI that takes options as argument
func NewEncodingInputsFtpCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsFtpCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsFtpCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsFtpCustomdataAPIWithClient constructor for EncodingInputsFtpCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsFtpCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsFtpCustomdataAPI {
    a := &EncodingInputsFtpCustomdataAPI{apiClient: apiClient}
    return a
}

// Get FTP Custom Data
func (api *EncodingInputsFtpCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/ftp/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

