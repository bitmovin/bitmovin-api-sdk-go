package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsSftpCustomdataAPI communicates with '/encoding/inputs/sftp/{input_id}/customData' endpoints
type EncodingInputsSftpCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsSftpCustomdataAPI constructor for EncodingInputsSftpCustomdataAPI that takes options as argument
func NewEncodingInputsSftpCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsSftpCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsSftpCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsSftpCustomdataAPIWithClient constructor for EncodingInputsSftpCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsSftpCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsSftpCustomdataAPI {
    a := &EncodingInputsSftpCustomdataAPI{apiClient: apiClient}
    return a
}

// Get SFTP Custom Data
func (api *EncodingInputsSftpCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/sftp/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

