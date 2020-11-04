package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsOutputsS3RoleBasedCustomdataAPI communicates with '/analytics/outputs/s3-role-based/{output_id}/customData' endpoints
type AnalyticsOutputsS3RoleBasedCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsOutputsS3RoleBasedCustomdataAPI constructor for AnalyticsOutputsS3RoleBasedCustomdataAPI that takes options as argument
func NewAnalyticsOutputsS3RoleBasedCustomdataAPI(options ...apiclient.APIClientOption) (*AnalyticsOutputsS3RoleBasedCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsOutputsS3RoleBasedCustomdataAPIWithClient(apiClient), nil
}

// NewAnalyticsOutputsS3RoleBasedCustomdataAPIWithClient constructor for AnalyticsOutputsS3RoleBasedCustomdataAPI that takes an APIClient as argument
func NewAnalyticsOutputsS3RoleBasedCustomdataAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsOutputsS3RoleBasedCustomdataAPI {
	a := &AnalyticsOutputsS3RoleBasedCustomdataAPI{apiClient: apiClient}
	return a
}

// Get S3 Role-based Output Custom Data
func (api *AnalyticsOutputsS3RoleBasedCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/analytics/outputs/s3-role-based/{output_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
