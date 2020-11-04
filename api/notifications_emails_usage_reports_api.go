package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsEmailsUsageReportsAPI communicates with '/notifications/emails/usage-reports' endpoints
type NotificationsEmailsUsageReportsAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsEmailsUsageReportsAPI constructor for NotificationsEmailsUsageReportsAPI that takes options as argument
func NewNotificationsEmailsUsageReportsAPI(options ...apiclient.APIClientOption) (*NotificationsEmailsUsageReportsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsEmailsUsageReportsAPIWithClient(apiClient), nil
}

// NewNotificationsEmailsUsageReportsAPIWithClient constructor for NotificationsEmailsUsageReportsAPI that takes an APIClient as argument
func NewNotificationsEmailsUsageReportsAPIWithClient(apiClient *apiclient.APIClient) *NotificationsEmailsUsageReportsAPI {
	a := &NotificationsEmailsUsageReportsAPI{apiClient: apiClient}
	return a
}

// List Email Notifications (All Usage Reports)
func (api *NotificationsEmailsUsageReportsAPI) List(queryParams ...func(*NotificationsEmailsUsageReportsAPIListQueryParams)) (*pagination.EmailNotificationsListPagination, error) {
	queryParameters := &NotificationsEmailsUsageReportsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EmailNotificationsListPagination
	err := api.apiClient.Get("/notifications/emails/usage-reports", nil, &responseModel, reqParams)
	return &responseModel, err
}

// NotificationsEmailsUsageReportsAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsEmailsUsageReportsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsEmailsUsageReportsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
