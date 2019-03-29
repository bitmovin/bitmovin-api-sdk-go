package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsWebhooksEncodingEncodingsFinishedApi struct {
    apiClient *common.ApiClient
    Customdata *NotificationsWebhooksEncodingEncodingsFinishedCustomdataApi
}

func NewNotificationsWebhooksEncodingEncodingsFinishedApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingEncodingsFinishedApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingEncodingsFinishedApi{apiClient: apiClient}

    customdataApi, err := NewNotificationsWebhooksEncodingEncodingsFinishedCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsWebhooksEncodingEncodingsFinishedApi) GetByWebhookId(webhookId string) (*model.Webhook, error) {
    var resp *model.Webhook
    reqParams := func(params *common.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
	}
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/finished/{webhook_id}", &resp, reqParams)
    return resp, err
}
func (api *NotificationsWebhooksEncodingEncodingsFinishedApi) GetByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.Webhook, error) {
    var resp *model.Webhook
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
	}
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/finished/{webhook_id}", &resp, reqParams)
    return resp, err
}
func (api *NotificationsWebhooksEncodingEncodingsFinishedApi) DeleteByWebhookId(webhookId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
	}
    err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/finished/{webhook_id}", &resp, reqParams)
    return resp, err
}
func (api *NotificationsWebhooksEncodingEncodingsFinishedApi) Create(webhook model.Webhook) (*model.Webhook, error) {
    payload := model.Webhook(webhook)
    
    err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/finished", &payload)
    return &payload, err
}
func (api *NotificationsWebhooksEncodingEncodingsFinishedApi) ListByEncodingId(encodingId string, queryParams ...func(*query.WebhookListByEncodingIdQueryParams)) (*pagination.WebhooksListByEncodingIdPagination, error) {
    queryParameters := &query.WebhookListByEncodingIdQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.WebhooksListByEncodingIdPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/finished", &resp, reqParams)
    return resp, err
}
func (api *NotificationsWebhooksEncodingEncodingsFinishedApi) CreateByEncodingId(encodingId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.Webhook(webhook)
    
    err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/finished", &payload, reqParams)
    return &payload, err
}
func (api *NotificationsWebhooksEncodingEncodingsFinishedApi) List(queryParams ...func(*query.WebhookListQueryParams)) (*pagination.WebhooksListPagination, error) {
    queryParameters := &query.WebhookListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.WebhooksListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/finished", &resp, reqParams)
    return resp, err
}
func (api *NotificationsWebhooksEncodingEncodingsFinishedApi) DeleteByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
	}
    err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/{encoding_id}/finished/{webhook_id}", &resp, reqParams)
    return resp, err
}
