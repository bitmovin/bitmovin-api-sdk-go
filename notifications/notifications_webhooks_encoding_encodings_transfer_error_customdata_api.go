package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsWebhooksEncodingEncodingsTransferErrorCustomdataApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataApi) GetCustomDataByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error/{webhook_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataApi) GetCustomDataByWebhookId(webhookId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/transfer-error/{webhook_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

