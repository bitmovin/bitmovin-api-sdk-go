package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp4InformationAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/information' endpoints
type EncodingEncodingsMuxingsMp4InformationAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp4InformationAPI constructor for EncodingEncodingsMuxingsMp4InformationAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4InformationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4InformationAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4InformationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4InformationAPIWithClient constructor for EncodingEncodingsMuxingsMp4InformationAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4InformationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4InformationAPI {
    a := &EncodingEncodingsMuxingsMp4InformationAPI{apiClient: apiClient}
    return a
}

// Get MP4 muxing Information
func (api *EncodingEncodingsMuxingsMp4InformationAPI) Get(encodingId string, muxingId string) (*model.Mp4MuxingInformation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.Mp4MuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/information", nil, &responseModel, reqParams)
    return &responseModel, err
}

