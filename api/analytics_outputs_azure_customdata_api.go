package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsOutputsAzureCustomdataAPI communicates with '/analytics/outputs/azure/{output_id}/customData' endpoints
type AnalyticsOutputsAzureCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsOutputsAzureCustomdataAPI constructor for AnalyticsOutputsAzureCustomdataAPI that takes options as argument
func NewAnalyticsOutputsAzureCustomdataAPI(options ...apiclient.APIClientOption) (*AnalyticsOutputsAzureCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsOutputsAzureCustomdataAPIWithClient(apiClient), nil
}

// NewAnalyticsOutputsAzureCustomdataAPIWithClient constructor for AnalyticsOutputsAzureCustomdataAPI that takes an APIClient as argument
func NewAnalyticsOutputsAzureCustomdataAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsOutputsAzureCustomdataAPI {
	a := &AnalyticsOutputsAzureCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Azure Output Custom Data
func (api *AnalyticsOutputsAzureCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/analytics/outputs/azure/{output_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
