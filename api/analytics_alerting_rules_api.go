package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AnalyticsAlertingRulesAPI communicates with '/analytics/alerting/rules' endpoints
type AnalyticsAlertingRulesAPI struct {
    apiClient *apiclient.APIClient

    // Threshold communicates with '/analytics/alerting/rules/{license_key}/threshold' endpoints
    Threshold *AnalyticsAlertingRulesThresholdAPI

}

// NewAnalyticsAlertingRulesAPI constructor for AnalyticsAlertingRulesAPI that takes options as argument
func NewAnalyticsAlertingRulesAPI(options ...apiclient.APIClientOption) (*AnalyticsAlertingRulesAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsAlertingRulesAPIWithClient(apiClient), nil
}

// NewAnalyticsAlertingRulesAPIWithClient constructor for AnalyticsAlertingRulesAPI that takes an APIClient as argument
func NewAnalyticsAlertingRulesAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAlertingRulesAPI {
    a := &AnalyticsAlertingRulesAPI{apiClient: apiClient}
    a.Threshold = NewAnalyticsAlertingRulesThresholdAPIWithClient(apiClient)

    return a
}

// Delete Analytics Alerting Rule By License Key And Id
func (api *AnalyticsAlertingRulesAPI) Delete(licenseKey string, ruleId string) (*model.AnalyticsAlertingRule, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["license_key"] = licenseKey
        params.PathParams["rule_id"] = ruleId
    }

    var responseModel model.AnalyticsAlertingRule
    err := api.apiClient.Delete("/analytics/alerting/rules/{license_key}/{rule_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Analytics Alerting Rule By License Key And Id
func (api *AnalyticsAlertingRulesAPI) Get(licenseKey string, ruleId string) (*model.AnalyticsAlertingRule, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["license_key"] = licenseKey
        params.PathParams["rule_id"] = ruleId
    }

    var responseModel model.AnalyticsAlertingRule
    err := api.apiClient.Get("/analytics/alerting/rules/{license_key}/{rule_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Analytics Alerting Rules
func (api *AnalyticsAlertingRulesAPI) List(queryParams ...func(*AnalyticsAlertingRulesAPIListQueryParams)) (*pagination.AnalyticsAlertingRulesListPagination, error) {
    queryParameters := &AnalyticsAlertingRulesAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.AnalyticsAlertingRulesListPagination
    err := api.apiClient.Get("/analytics/alerting/rules", nil, &responseModel, reqParams)
    return &responseModel, err
}

// AnalyticsAlertingRulesAPIListQueryParams contains all query parameters for the List endpoint
type AnalyticsAlertingRulesAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *AnalyticsAlertingRulesAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


