package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsSmoothCustomdataAPI communicates with '/encoding/manifests/smooth/{manifest_id}/customData' endpoints
type EncodingManifestsSmoothCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsSmoothCustomdataAPI constructor for EncodingManifestsSmoothCustomdataAPI that takes options as argument
func NewEncodingManifestsSmoothCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingManifestsSmoothCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsSmoothCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingManifestsSmoothCustomdataAPIWithClient constructor for EncodingManifestsSmoothCustomdataAPI that takes an APIClient as argument
func NewEncodingManifestsSmoothCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsSmoothCustomdataAPI {
    a := &EncodingManifestsSmoothCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Smooth Streaming Manifest Custom Data
func (api *EncodingManifestsSmoothCustomdataAPI) Get(manifestId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

