package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AnalyticsOutputsAPI intermediary API object with no endpoints
type AnalyticsOutputsAPI struct {
	apiClient *apiclient.APIClient

	// Azure communicates with '/analytics/outputs/azure' endpoints
	Azure *AnalyticsOutputsAzureAPI
	// S3RoleBased communicates with '/analytics/outputs/s3-role-based' endpoints
	S3RoleBased *AnalyticsOutputsS3RoleBasedAPI
	// GcsServiceAccount communicates with '/analytics/outputs/gcs-service-account' endpoints
	GcsServiceAccount *AnalyticsOutputsGcsServiceAccountAPI
}

// NewAnalyticsOutputsAPI constructor for AnalyticsOutputsAPI that takes options as argument
func NewAnalyticsOutputsAPI(options ...apiclient.APIClientOption) (*AnalyticsOutputsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsOutputsAPIWithClient(apiClient), nil
}

// NewAnalyticsOutputsAPIWithClient constructor for AnalyticsOutputsAPI that takes an APIClient as argument
func NewAnalyticsOutputsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsOutputsAPI {
	a := &AnalyticsOutputsAPI{apiClient: apiClient}
	a.Azure = NewAnalyticsOutputsAzureAPIWithClient(apiClient)
	a.S3RoleBased = NewAnalyticsOutputsS3RoleBasedAPIWithClient(apiClient)
	a.GcsServiceAccount = NewAnalyticsOutputsGcsServiceAccountAPIWithClient(apiClient)

	return a
}
