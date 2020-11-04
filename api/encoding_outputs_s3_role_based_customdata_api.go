package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsS3RoleBasedCustomdataAPI communicates with '/encoding/outputs/s3-role-based/{output_id}/customData' endpoints
type EncodingOutputsS3RoleBasedCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingOutputsS3RoleBasedCustomdataAPI constructor for EncodingOutputsS3RoleBasedCustomdataAPI that takes options as argument
func NewEncodingOutputsS3RoleBasedCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsS3RoleBasedCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsS3RoleBasedCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsS3RoleBasedCustomdataAPIWithClient constructor for EncodingOutputsS3RoleBasedCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsS3RoleBasedCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsS3RoleBasedCustomdataAPI {
	a := &EncodingOutputsS3RoleBasedCustomdataAPI{apiClient: apiClient}
	return a
}

// Get S3 Role-based Output Custom Data
func (api *EncodingOutputsS3RoleBasedCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/outputs/s3-role-based/{output_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
