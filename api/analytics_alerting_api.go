package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AnalyticsAlertingAPI intermediary API object with no endpoints
type AnalyticsAlertingAPI struct {
	apiClient *apiclient.APIClient

	// Rules communicates with '/analytics/alerting/rules' endpoints
	Rules *AnalyticsAlertingRulesAPI
	// Incidents communicates with '/analytics/alerting/incidents' endpoints
	Incidents *AnalyticsAlertingIncidentsAPI
}

// NewAnalyticsAlertingAPI constructor for AnalyticsAlertingAPI that takes options as argument
func NewAnalyticsAlertingAPI(options ...apiclient.APIClientOption) (*AnalyticsAlertingAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsAlertingAPIWithClient(apiClient), nil
}

// NewAnalyticsAlertingAPIWithClient constructor for AnalyticsAlertingAPI that takes an APIClient as argument
func NewAnalyticsAlertingAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAlertingAPI {
	a := &AnalyticsAlertingAPI{apiClient: apiClient}
	a.Rules = NewAnalyticsAlertingRulesAPIWithClient(apiClient)
	a.Incidents = NewAnalyticsAlertingIncidentsAPIWithClient(apiClient)

	return a
}
