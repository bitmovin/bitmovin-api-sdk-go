package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveMovInformationAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}/information' endpoints
type EncodingEncodingsMuxingsProgressiveMovInformationAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveMovInformationAPI constructor for EncodingEncodingsMuxingsProgressiveMovInformationAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveMovInformationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveMovInformationAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveMovInformationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveMovInformationAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveMovInformationAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveMovInformationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveMovInformationAPI {
    a := &EncodingEncodingsMuxingsProgressiveMovInformationAPI{apiClient: apiClient}
    return a
}

// Get Progressive MOV muxing Information
func (api *EncodingEncodingsMuxingsProgressiveMovInformationAPI) Get(encodingId string, muxingId string) (*model.ProgressiveMovMuxingInformation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.ProgressiveMovMuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}/information", nil, &responseModel, reqParams)
    return &responseModel, err
}

