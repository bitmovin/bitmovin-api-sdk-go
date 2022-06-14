package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsDirectFileUploadCustomdataAPI communicates with '/encoding/inputs/direct-file-upload/{input_id}/customData' endpoints
type EncodingInputsDirectFileUploadCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsDirectFileUploadCustomdataAPI constructor for EncodingInputsDirectFileUploadCustomdataAPI that takes options as argument
func NewEncodingInputsDirectFileUploadCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsDirectFileUploadCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsDirectFileUploadCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsDirectFileUploadCustomdataAPIWithClient constructor for EncodingInputsDirectFileUploadCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsDirectFileUploadCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsDirectFileUploadCustomdataAPI {
	a := &EncodingInputsDirectFileUploadCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Direct File Upload Custom Data
func (api *EncodingInputsDirectFileUploadCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/inputs/direct-file-upload/{input_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
