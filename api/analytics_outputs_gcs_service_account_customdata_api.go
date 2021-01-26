package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsOutputsGcsServiceAccountCustomdataAPI communicates with '/analytics/outputs/gcs-service-account/{output_id}/customData' endpoints
type AnalyticsOutputsGcsServiceAccountCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsOutputsGcsServiceAccountCustomdataAPI constructor for AnalyticsOutputsGcsServiceAccountCustomdataAPI that takes options as argument
func NewAnalyticsOutputsGcsServiceAccountCustomdataAPI(options ...apiclient.APIClientOption) (*AnalyticsOutputsGcsServiceAccountCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsOutputsGcsServiceAccountCustomdataAPIWithClient(apiClient), nil
}

// NewAnalyticsOutputsGcsServiceAccountCustomdataAPIWithClient constructor for AnalyticsOutputsGcsServiceAccountCustomdataAPI that takes an APIClient as argument
func NewAnalyticsOutputsGcsServiceAccountCustomdataAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsOutputsGcsServiceAccountCustomdataAPI {
	a := &AnalyticsOutputsGcsServiceAccountCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Service Account based GCS Output Custom Data
func (api *AnalyticsOutputsGcsServiceAccountCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/analytics/outputs/gcs-service-account/{output_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
