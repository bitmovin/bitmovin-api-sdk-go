package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI communicates with '/notifications/webhooks/encoding/encodings/live-encoding-heartbeat' endpoints
type NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI constructor for NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI {
	a := &NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI{apiClient: apiClient}
	return a
}

// Create Add &#39;Live Encoding Heartbeat&#39; Webhook
// Add a new webhook notification that triggers a heartbeat webhook with a fixed &#x60;interval&#x60; for all Live Encodings.
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI) Create(liveEncodingHeartbeatWebhook model.LiveEncodingHeartbeatWebhook) (*model.LiveEncodingHeartbeatWebhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.LiveEncodingHeartbeatWebhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/live-encoding-heartbeat", &liveEncodingHeartbeatWebhook, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteByWebhookId Delete &#39;Live Encoding Heartbeat&#39; Webhook
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI) DeleteByWebhookId(webhookId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/live-encoding-heartbeat/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetByWebhookId &#39;Live Encoding Heartbeat&#39; Webhook Details
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI) GetByWebhookId(webhookId string) (*model.LiveEncodingHeartbeatWebhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.LiveEncodingHeartbeatWebhook
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/live-encoding-heartbeat/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List &#39;Live Encoding Heartbeat&#39; Webhooks
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI) List(queryParams ...func(*NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListQueryParams)) (*pagination.LiveEncodingHeartbeatWebhooksListPagination, error) {
	queryParameters := &NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.LiveEncodingHeartbeatWebhooksListPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/live-encoding-heartbeat", nil, &responseModel, reqParams)
	return &responseModel, err
}

// NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
