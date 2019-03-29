package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsStateApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsStateApi(configs ...func(*common.ApiClient)) (*NotificationsStateApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsStateApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsStateApi) List(notificationId string, resourceId string, queryParams ...func(*query.NotificationStateEntryListQueryParams)) (*pagination.NotificationStateEntrysListPagination, error) {
    queryParameters := &query.NotificationStateEntryListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.NotificationStateEntrysListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
        params.PathParams["resource_id"] = resourceId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/notifications/{notification_id}/state/{resource_id}", &resp, reqParams)
    return resp, err
}
