package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAlertingRulesThresholdAPI communicates with '/analytics/alerting/rules/{license_key}/threshold' endpoints
type AnalyticsAlertingRulesThresholdAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsAlertingRulesThresholdAPI constructor for AnalyticsAlertingRulesThresholdAPI that takes options as argument
func NewAnalyticsAlertingRulesThresholdAPI(options ...apiclient.APIClientOption) (*AnalyticsAlertingRulesThresholdAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsAlertingRulesThresholdAPIWithClient(apiClient), nil
}

// NewAnalyticsAlertingRulesThresholdAPIWithClient constructor for AnalyticsAlertingRulesThresholdAPI that takes an APIClient as argument
func NewAnalyticsAlertingRulesThresholdAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAlertingRulesThresholdAPI {
	a := &AnalyticsAlertingRulesThresholdAPI{apiClient: apiClient}
	return a
}

// Create Analytics Alerting Rule
func (api *AnalyticsAlertingRulesThresholdAPI) Create(licenseKey string, analyticsAlertingRule model.AnalyticsAlertingRule) (*model.AnalyticsAlertingRule, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_key"] = licenseKey
	}

	var responseModel model.AnalyticsAlertingRule
	err := api.apiClient.Post("/analytics/alerting/rules/{license_key}/threshold", &analyticsAlertingRule, &responseModel, reqParams)
	return &responseModel, err
}

// Update Analytics Alerting Rule
func (api *AnalyticsAlertingRulesThresholdAPI) Update(licenseKey string, ruleId string, analyticsAlertingRule model.AnalyticsAlertingRule) (*model.AnalyticsAlertingRule, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_key"] = licenseKey
		params.PathParams["rule_id"] = ruleId
	}

	var responseModel model.AnalyticsAlertingRule
	err := api.apiClient.Put("/analytics/alerting/rules/{license_key}/threshold/{rule_id}", &analyticsAlertingRule, &responseModel, reqParams)
	return &responseModel, err
}
