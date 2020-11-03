package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsCmafCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsCmafCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsCmafCustomdataAPI constructor for EncodingEncodingsMuxingsCmafCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsCmafCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsCmafCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsCmafCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsCmafCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsCmafCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsCmafCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsCmafCustomdataAPI {
    a := &EncodingEncodingsMuxingsCmafCustomdataAPI{apiClient: apiClient}
    return a
}

// Get CMAF muxing custom data
func (api *EncodingEncodingsMuxingsCmafCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

