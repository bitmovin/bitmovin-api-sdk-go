package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsWebhooksEncodingEncodingsTransferErrorApi struct {
    apiClient *common.ApiClient
    Customdata *NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataApi
}

func NewNotificationsWebhooksEncodingEncodingsTransferErrorApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingEncodingsTransferErrorApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingEncodingsTransferErrorApi{apiClient: apiClient}

    customdataApi, err := NewNotificationsWebhooksEncodingEncodingsTransferErrorCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorApi) Create(webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/transfer-error", &webhook, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorApi) CreateByEncodingId(encodingId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error", &webhook, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorApi) DeleteByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error/{webhook_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorApi) DeleteByWebhookId(webhookId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/transfer-error/{webhook_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorApi) GetByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error/{webhook_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorApi) GetByWebhookId(webhookId string) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/transfer-error/{webhook_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorApi) List(queryParams ...func(*query.WebhookListQueryParams)) (*pagination.WebhooksListPagination, error) {
    queryParameters := &query.WebhookListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.WebhooksListPagination
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/transfer-error", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorApi) ListByEncodingId(encodingId string, queryParams ...func(*query.WebhookListByEncodingIdQueryParams)) (*pagination.WebhooksListByEncodingIdPagination, error) {
    queryParameters := &query.WebhookListByEncodingIdQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.WebhooksListByEncodingIdPagination
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error", nil, &responseModel, reqParams)
    return responseModel, err
}

