package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsMoqCustomdataAPI communicates with '/encoding/outputs/moq/{output_id}/customData' endpoints
type EncodingOutputsMoqCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingOutputsMoqCustomdataAPI constructor for EncodingOutputsMoqCustomdataAPI that takes options as argument
func NewEncodingOutputsMoqCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsMoqCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsMoqCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsMoqCustomdataAPIWithClient constructor for EncodingOutputsMoqCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsMoqCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsMoqCustomdataAPI {
	a := &EncodingOutputsMoqCustomdataAPI{apiClient: apiClient}
	return a
}

// Get MoQ Output Custom Data
func (api *EncodingOutputsMoqCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/outputs/moq/{output_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
