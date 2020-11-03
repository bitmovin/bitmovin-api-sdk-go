package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersDenoiseHqdn3dCustomdataAPI communicates with '/encoding/filters/denoise-hqdn3d/{filter_id}/customData' endpoints
type EncodingFiltersDenoiseHqdn3dCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingFiltersDenoiseHqdn3dCustomdataAPI constructor for EncodingFiltersDenoiseHqdn3dCustomdataAPI that takes options as argument
func NewEncodingFiltersDenoiseHqdn3dCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersDenoiseHqdn3dCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersDenoiseHqdn3dCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersDenoiseHqdn3dCustomdataAPIWithClient constructor for EncodingFiltersDenoiseHqdn3dCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersDenoiseHqdn3dCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersDenoiseHqdn3dCustomdataAPI {
    a := &EncodingFiltersDenoiseHqdn3dCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Denoise hqdn3d Filter Custom Data
func (api *EncodingFiltersDenoiseHqdn3dCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/filters/denoise-hqdn3d/{filter_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

