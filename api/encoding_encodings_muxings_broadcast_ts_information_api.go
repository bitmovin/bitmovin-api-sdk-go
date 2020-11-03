package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsBroadcastTsInformationAPI communicates with '/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}/information' endpoints
type EncodingEncodingsMuxingsBroadcastTsInformationAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsBroadcastTsInformationAPI constructor for EncodingEncodingsMuxingsBroadcastTsInformationAPI that takes options as argument
func NewEncodingEncodingsMuxingsBroadcastTsInformationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsBroadcastTsInformationAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsBroadcastTsInformationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsBroadcastTsInformationAPIWithClient constructor for EncodingEncodingsMuxingsBroadcastTsInformationAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsBroadcastTsInformationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsBroadcastTsInformationAPI {
    a := &EncodingEncodingsMuxingsBroadcastTsInformationAPI{apiClient: apiClient}
    return a
}

// Get Broadcast TS muxing Information
func (api *EncodingEncodingsMuxingsBroadcastTsInformationAPI) Get(encodingId string, muxingId string) (*model.BroadcastTsMuxingInformation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.BroadcastTsMuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}/information", nil, &responseModel, reqParams)
    return &responseModel, err
}

