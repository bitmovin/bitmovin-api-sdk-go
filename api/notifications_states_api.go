package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsStatesAPI communicates with '/notifications/{notification_id}/states/{resource_id}' endpoints
type NotificationsStatesAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsStatesAPI constructor for NotificationsStatesAPI that takes options as argument
func NewNotificationsStatesAPI(options ...apiclient.APIClientOption) (*NotificationsStatesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsStatesAPIWithClient(apiClient), nil
}

// NewNotificationsStatesAPIWithClient constructor for NotificationsStatesAPI that takes an APIClient as argument
func NewNotificationsStatesAPIWithClient(apiClient *apiclient.APIClient) *NotificationsStatesAPI {
	a := &NotificationsStatesAPI{apiClient: apiClient}
	return a
}

// List Notification State History (Specific Resource)
func (api *NotificationsStatesAPI) List(notificationId string, resourceId string, queryParams ...func(*NotificationsStatesAPIListQueryParams)) (*pagination.NotificationStateEntrysListPagination, error) {
	queryParameters := &NotificationsStatesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["notification_id"] = notificationId
		params.PathParams["resource_id"] = resourceId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.NotificationStateEntrysListPagination
	err := api.apiClient.Get("/notifications/{notification_id}/states/{resource_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// NotificationsStatesAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsStatesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsStatesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
