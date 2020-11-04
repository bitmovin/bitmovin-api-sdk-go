package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsAkamaiMslCustomdataAPI communicates with '/encoding/outputs/akamai-msl/{output_id}/customData' endpoints
type EncodingOutputsAkamaiMslCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingOutputsAkamaiMslCustomdataAPI constructor for EncodingOutputsAkamaiMslCustomdataAPI that takes options as argument
func NewEncodingOutputsAkamaiMslCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsAkamaiMslCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsAkamaiMslCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsAkamaiMslCustomdataAPIWithClient constructor for EncodingOutputsAkamaiMslCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsAkamaiMslCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsAkamaiMslCustomdataAPI {
	a := &EncodingOutputsAkamaiMslCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Akamai MSL Output Custom Data
func (api *EncodingOutputsAkamaiMslCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/outputs/akamai-msl/{output_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
