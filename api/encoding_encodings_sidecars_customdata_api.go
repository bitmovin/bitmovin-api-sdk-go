package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsSidecarsCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}/customData' endpoints
type EncodingEncodingsSidecarsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsSidecarsCustomdataAPI constructor for EncodingEncodingsSidecarsCustomdataAPI that takes options as argument
func NewEncodingEncodingsSidecarsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsSidecarsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsSidecarsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsSidecarsCustomdataAPIWithClient constructor for EncodingEncodingsSidecarsCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsSidecarsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsSidecarsCustomdataAPI {
    a := &EncodingEncodingsSidecarsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Sidecar Custom Data
func (api *EncodingEncodingsSidecarsCustomdataAPI) Get(encodingId string, sidecarId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["sidecar_id"] = sidecarId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

