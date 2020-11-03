package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsHlsCustomdataAPI communicates with '/encoding/manifests/hls/{manifest_id}/customData' endpoints
type EncodingManifestsHlsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsCustomdataAPI constructor for EncodingManifestsHlsCustomdataAPI that takes options as argument
func NewEncodingManifestsHlsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsHlsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsCustomdataAPIWithClient constructor for EncodingManifestsHlsCustomdataAPI that takes an APIClient as argument
func NewEncodingManifestsHlsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsCustomdataAPI {
    a := &EncodingManifestsHlsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get HLS Manifest Custom Data
func (api *EncodingManifestsHlsCustomdataAPI) Get(manifestId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

