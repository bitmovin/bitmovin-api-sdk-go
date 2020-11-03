package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsDashCustomdataAPI communicates with '/encoding/manifests/dash/{manifest_id}/customData' endpoints
type EncodingManifestsDashCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashCustomdataAPI constructor for EncodingManifestsDashCustomdataAPI that takes options as argument
func NewEncodingManifestsDashCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashCustomdataAPIWithClient constructor for EncodingManifestsDashCustomdataAPI that takes an APIClient as argument
func NewEncodingManifestsDashCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashCustomdataAPI {
    a := &EncodingManifestsDashCustomdataAPI{apiClient: apiClient}
    return a
}

// Get DASH Manifest Custom Data
func (api *EncodingManifestsDashCustomdataAPI) Get(manifestId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

