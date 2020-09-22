package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsTypeAPI communicates with '/encoding/inputs/{input_id}/type' endpoints
type EncodingInputsTypeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsTypeAPI constructor for EncodingInputsTypeAPI that takes options as argument
func NewEncodingInputsTypeAPI(options ...apiclient.APIClientOption) (*EncodingInputsTypeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsTypeAPIWithClient(apiClient), nil
}

// NewEncodingInputsTypeAPIWithClient constructor for EncodingInputsTypeAPI that takes an APIClient as argument
func NewEncodingInputsTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsTypeAPI {
	a := &EncodingInputsTypeAPI{apiClient: apiClient}
	return a
}

// Get Input Type
func (api *EncodingInputsTypeAPI) Get(inputId string) (*model.InputTypeResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.InputTypeResponse
	err := api.apiClient.Get("/encoding/inputs/{input_id}/type", nil, &responseModel, reqParams)
	return &responseModel, err
}
