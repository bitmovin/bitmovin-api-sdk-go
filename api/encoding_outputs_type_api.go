package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsTypeAPI communicates with '/encoding/outputs/{output_id}/type' endpoints
type EncodingOutputsTypeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingOutputsTypeAPI constructor for EncodingOutputsTypeAPI that takes options as argument
func NewEncodingOutputsTypeAPI(options ...apiclient.APIClientOption) (*EncodingOutputsTypeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsTypeAPIWithClient(apiClient), nil
}

// NewEncodingOutputsTypeAPIWithClient constructor for EncodingOutputsTypeAPI that takes an APIClient as argument
func NewEncodingOutputsTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsTypeAPI {
	a := &EncodingOutputsTypeAPI{apiClient: apiClient}
	return a
}

// Get Output Type
func (api *EncodingOutputsTypeAPI) Get(outputId string) (*model.OutputTypeResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.OutputTypeResponse
	err := api.apiClient.Get("/encoding/outputs/{output_id}/type", nil, &responseModel, reqParams)
	return &responseModel, err
}
