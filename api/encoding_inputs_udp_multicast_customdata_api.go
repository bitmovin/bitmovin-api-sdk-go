package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsUdpMulticastCustomdataAPI communicates with '/encoding/inputs/udp-multicast/{input_id}/customData' endpoints
type EncodingInputsUdpMulticastCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsUdpMulticastCustomdataAPI constructor for EncodingInputsUdpMulticastCustomdataAPI that takes options as argument
func NewEncodingInputsUdpMulticastCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsUdpMulticastCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsUdpMulticastCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsUdpMulticastCustomdataAPIWithClient constructor for EncodingInputsUdpMulticastCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsUdpMulticastCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsUdpMulticastCustomdataAPI {
    a := &EncodingInputsUdpMulticastCustomdataAPI{apiClient: apiClient}
    return a
}

// Get UDP multicast input Custom Data
func (api *EncodingInputsUdpMulticastCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/udp-multicast/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

