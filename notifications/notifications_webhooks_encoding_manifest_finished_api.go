package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type NotificationsWebhooksEncodingManifestFinishedApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsWebhooksEncodingManifestFinishedApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingManifestFinishedApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingManifestFinishedApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsWebhooksEncodingManifestFinishedApi) Create(webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/manifest/finished", &webhook, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingManifestFinishedApi) CreateByManifestId(manifestId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/manifest/{manifest_id}/finished", &webhook, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingManifestFinishedApi) Delete(notificationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/webhooks/encoding/manifest/finished/{notification_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingManifestFinishedApi) Update(notificationId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Put("/notifications/webhooks/encoding/manifest/finished/{notification_id}", &webhook, &responseModel, reqParams)
    return responseModel, err
}

