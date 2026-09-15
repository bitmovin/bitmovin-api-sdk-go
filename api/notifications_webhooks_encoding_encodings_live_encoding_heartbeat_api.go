package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI communicates with '/notifications/webhooks/encoding/encodings/live-encoding-heartbeat' endpoints
type NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/notifications/webhooks/encoding/encodings/live-encoding-heartbeat/{webhook_id}/customData' endpoints
	Customdata *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI
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
	a.Customdata = NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPIWithClient(apiClient)

	return a
}

// Create &#39;Live Encoding Heartbeat&#39; Webhook
// Add a new webhook notification that triggers a heartbeat webhook with a fixed &#x60;interval&#x60; for all Live Encodings.
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI) Create(liveEncodingHeartbeatWebhook model.LiveEncodingHeartbeatWebhook) (*model.LiveEncodingHeartbeatWebhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.LiveEncodingHeartbeatWebhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/live-encoding-heartbeat", &liveEncodingHeartbeatWebhook, &responseModel, reqParams)
	return &responseModel, err
}

// CreateByEncodingId Create &#39;Live Encoding Heartbeat&#39; Webhook for a specific Encoding
// Add a webhook notification that triggers a heartbeat with a fixed &#x60;interval&#x60; for a specific encoding. A maximum number of 5 webhooks per Encoding is allowed. Can also be added to an already running live encoding (on a supported encoder version); the first heartbeat fires immediately, then one every &#x60;interval&#x60; seconds.
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI) CreateByEncodingId(encodingId string, liveEncodingHeartbeatWebhook model.LiveEncodingHeartbeatWebhook) (*model.LiveEncodingHeartbeatWebhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.LiveEncodingHeartbeatWebhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/live-encoding-heartbeat", &liveEncodingHeartbeatWebhook, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteByEncodingIdAndWebhookId Delete &#39;Live Encoding Heartbeat&#39; Webhook for a specific Encoding
// Delete a &#39;Live Encoding Heartbeat&#39; webhook for a specific encoding. Can also be deleted while the live encoding is running (on a supported encoder version); no further heartbeats are sent, but deliveries already in flight are not cancelled.
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI) DeleteByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/{encoding_id}/live-encoding-heartbeat/{webhook_id}", nil, &responseModel, reqParams)
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

// GetByEncodingIdAndWebhookId Get &#39;Live Encoding Heartbeat&#39; Webhook details for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI) GetByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.LiveEncodingHeartbeatWebhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.LiveEncodingHeartbeatWebhook
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/live-encoding-heartbeat/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetByWebhookId Get &#39;Live Encoding Heartbeat&#39; Webhook details
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

// ListByEncodingId List &#39;Live Encoding Heartbeat&#39; Webhooks for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPI) ListByEncodingId(encodingId string, queryParams ...func(*NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListByEncodingIdQueryParams)) (*pagination.LiveEncodingHeartbeatWebhooksListByEncodingIdPagination, error) {
	queryParameters := &NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListByEncodingIdQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.LiveEncodingHeartbeatWebhooksListByEncodingIdPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/live-encoding-heartbeat", nil, &responseModel, reqParams)
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

// NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListByEncodingIdQueryParams contains all query parameters for the ListByEncodingId endpoint
type NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListByEncodingIdQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatAPIListByEncodingIdQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
