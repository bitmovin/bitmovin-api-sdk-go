package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI communicates with '/notifications/webhooks/encoding/encodings/live-manifest-ready' endpoints
type NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/notifications/webhooks/encoding/encodings/live-manifest-ready/{webhook_id}/customData' endpoints
	Customdata *NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI
}

// NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI constructor for NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI {
	a := &NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI{apiClient: apiClient}
	a.Customdata = NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add &#39;Live Manifest Ready&#39; Webhook
// Add a webhook notification that is triggered when a live manifest is ready, meaning at least one segment is available in every playlist. A maximum number of 5 webhooks is allowed
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI) Create(webhook model.Webhook) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Webhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/live-manifest-ready", &webhook, &responseModel, reqParams)
	return &responseModel, err
}

// CreateByEncodingId Add &#39;Live Manifest Ready&#39; Webhook for a specific Encoding
// Add a webhook notification that triggers when a live manifest is ready for a specific encoding. A maximum number of 5 webhooks per Encoding is allowed. For a running live encoding (V2 manifest generator, encoder version &#x60;2.276.0&#x60; or above) it is propagated asynchronously (eventually consistent, typically sub-second) and does not require a restart; if the manifest is already ready, it fires immediately and exactly once. Org-level webhooks apply on the next start.
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI) CreateByEncodingId(encodingId string, webhook model.Webhook) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Webhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/live-manifest-ready", &webhook, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteByEncodingIdAndWebhookId Delete &#39;Live Manifest Ready&#39; Webhook for a specific Encoding
// Delete a &#39;Live Manifest Ready&#39; webhook for a specific encoding. For a running live encoding (V2 manifest generator, encoder version &#x60;2.276.0&#x60; or above) it is propagated asynchronously (eventually consistent, typically sub-second) and does not require a restart; deliveries already in flight are not cancelled. Org-level webhooks apply on the next start.
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI) DeleteByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/{encoding_id}/live-manifest-ready/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteByWebhookId Delete &#39;Live Manifest Ready&#39; Webhook
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI) DeleteByWebhookId(webhookId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/live-manifest-ready/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetByEncodingIdAndWebhookId &#39;Live Manifest Ready&#39; Webhook Details for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI) GetByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.Webhook
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/live-manifest-ready/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetByWebhookId &#39;Live Manifest Ready&#39; Webhook Details
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI) GetByWebhookId(webhookId string) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.Webhook
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/live-manifest-ready/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List &#39;Live Manifest Ready&#39; Webhooks
// List all Live Manifest Ready webhook notifications. A maximum number of 5 webhooks is allowed
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI) List(queryParams ...func(*NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListQueryParams)) (*pagination.WebhooksListPagination, error) {
	queryParameters := &NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WebhooksListPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/live-manifest-ready", nil, &responseModel, reqParams)
	return &responseModel, err
}

// ListByEncodingId List &#39;Live Manifest Ready&#39; Webhooks for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPI) ListByEncodingId(encodingId string, queryParams ...func(*NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListByEncodingIdQueryParams)) (*pagination.WebhooksListByEncodingIdPagination, error) {
	queryParameters := &NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListByEncodingIdQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WebhooksListByEncodingIdPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/live-manifest-ready", nil, &responseModel, reqParams)
	return &responseModel, err
}

// NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}

// NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListByEncodingIdQueryParams contains all query parameters for the ListByEncodingId endpoint
type NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListByEncodingIdQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsLiveManifestReadyAPIListByEncodingIdQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
