package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsEmailsEncodingApi struct {
    apiClient *common.ApiClient
    Encodings *NotificationsEmailsEncodingEncodingsApi
}

func NewNotificationsEmailsEncodingApi(configs ...func(*common.ApiClient)) (*NotificationsEmailsEncodingApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsEmailsEncodingApi{apiClient: apiClient}

    encodingsApi, err := NewNotificationsEmailsEncodingEncodingsApi(configs...)
    api.Encodings = encodingsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsEmailsEncodingApi) List(queryParams ...func(*query.EmailNotificationListQueryParams)) (*pagination.EmailNotificationsListPagination, error) {
    queryParameters := &query.EmailNotificationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.EmailNotificationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/notifications/emails/encoding", &resp, reqParams)
    return resp, err
}
