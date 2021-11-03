package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsWebhooksEncodingEncodingsTransferErrorAPI communicates with '/notifications/webhooks/encoding/encodings/transfer-error' endpoints
type NotificationsWebhooksEncodingEncodingsTransferErrorAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/notifications/webhooks/encoding/encodings/transfer-error/{webhook_id}/customData' endpoints
	Customdata *NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI
}

// NewNotificationsWebhooksEncodingEncodingsTransferErrorAPI constructor for NotificationsWebhooksEncodingEncodingsTransferErrorAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsTransferErrorAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsTransferErrorAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsTransferErrorAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsTransferErrorAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsTransferErrorAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsTransferErrorAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsTransferErrorAPI {
	a := &NotificationsWebhooksEncodingEncodingsTransferErrorAPI{apiClient: apiClient}
	a.Customdata = NewNotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add &#39;Encoding Transfer Error&#39; Webhook
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorAPI) Create(webhook model.Webhook) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Webhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/transfer-error", &webhook, &responseModel, reqParams)
	return &responseModel, err
}

// CreateByEncodingId Add &#39;Encoding Transfer Error&#39; Webhook for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorAPI) CreateByEncodingId(encodingId string, webhook model.Webhook) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Webhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error", &webhook, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteByEncodingIdAndWebhookId Delete &#39;Encoding Transfer Error&#39; Webhook for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorAPI) DeleteByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteByWebhookId Delete &#39;Encoding Transfer Error&#39; Webhook
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorAPI) DeleteByWebhookId(webhookId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/transfer-error/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetByEncodingIdAndWebhookId &#39;Encoding Transfer Error&#39; Webhook Details for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorAPI) GetByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.Webhook
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetByWebhookId &#39;Encoding Transfer Error&#39; Webhook Details
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorAPI) GetByWebhookId(webhookId string) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.Webhook
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/transfer-error/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List &#39;Encoding Transfer Error&#39; Webhooks
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorAPI) List(queryParams ...func(*NotificationsWebhooksEncodingEncodingsTransferErrorAPIListQueryParams)) (*pagination.WebhooksListPagination, error) {
	queryParameters := &NotificationsWebhooksEncodingEncodingsTransferErrorAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WebhooksListPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/transfer-error", nil, &responseModel, reqParams)
	return &responseModel, err
}

// ListByEncodingId List &#39;Encoding Transfer Error&#39; Webhooks for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorAPI) ListByEncodingId(encodingId string, queryParams ...func(*NotificationsWebhooksEncodingEncodingsTransferErrorAPIListByEncodingIdQueryParams)) (*pagination.WebhooksListByEncodingIdPagination, error) {
	queryParameters := &NotificationsWebhooksEncodingEncodingsTransferErrorAPIListByEncodingIdQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WebhooksListByEncodingIdPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error", nil, &responseModel, reqParams)
	return &responseModel, err
}

// NotificationsWebhooksEncodingEncodingsTransferErrorAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsWebhooksEncodingEncodingsTransferErrorAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsTransferErrorAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}

// NotificationsWebhooksEncodingEncodingsTransferErrorAPIListByEncodingIdQueryParams contains all query parameters for the ListByEncodingId endpoint
type NotificationsWebhooksEncodingEncodingsTransferErrorAPIListByEncodingIdQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsTransferErrorAPIListByEncodingIdQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
