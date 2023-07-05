package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsWebhooksEncodingEncodingsFinishedAPI communicates with '/notifications/webhooks/encoding/encodings/finished' endpoints
type NotificationsWebhooksEncodingEncodingsFinishedAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/notifications/webhooks/encoding/encodings/finished/{webhook_id}/customData' endpoints
	Customdata *NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI
}

// NewNotificationsWebhooksEncodingEncodingsFinishedAPI constructor for NotificationsWebhooksEncodingEncodingsFinishedAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsFinishedAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsFinishedAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsFinishedAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsFinishedAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsFinishedAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsFinishedAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsFinishedAPI {
	a := &NotificationsWebhooksEncodingEncodingsFinishedAPI{apiClient: apiClient}
	a.Customdata = NewNotificationsWebhooksEncodingEncodingsFinishedCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add &#39;Encoding Finished&#39; Webhook
// Add a new webhook notification that triggers if an encoding finishes successfully. A maximum number of 5 webhooks is allowed
func (api *NotificationsWebhooksEncodingEncodingsFinishedAPI) Create(webhook model.Webhook) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Webhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/finished", &webhook, &responseModel, reqParams)
	return &responseModel, err
}

// CreateByEncodingId Add &#39;Encoding Finished&#39; Webhook for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsFinishedAPI) CreateByEncodingId(encodingId string, webhook model.Webhook) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Webhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/finished", &webhook, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteByEncodingIdAndWebhookId Delete &#39;Encoding Finished&#39; Webhook for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsFinishedAPI) DeleteByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/{encoding_id}/finished/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteByWebhookId Delete &#39;Encoding Finished&#39; Webhook
func (api *NotificationsWebhooksEncodingEncodingsFinishedAPI) DeleteByWebhookId(webhookId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/finished/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetByEncodingIdAndWebhookId &#39;Encoding Finished&#39; Webhook Details for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsFinishedAPI) GetByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.Webhook
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/finished/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetByWebhookId &#39;Encoding Finished&#39; Webhook Details
func (api *NotificationsWebhooksEncodingEncodingsFinishedAPI) GetByWebhookId(webhookId string) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.Webhook
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/finished/{webhook_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List &#39;Encoding Finished&#39; Webhooks
func (api *NotificationsWebhooksEncodingEncodingsFinishedAPI) List(queryParams ...func(*NotificationsWebhooksEncodingEncodingsFinishedAPIListQueryParams)) (*pagination.WebhooksListPagination, error) {
	queryParameters := &NotificationsWebhooksEncodingEncodingsFinishedAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WebhooksListPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/finished", nil, &responseModel, reqParams)
	return &responseModel, err
}

// ListByEncodingId List &#39;Encoding Finished&#39; Webhooks for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsFinishedAPI) ListByEncodingId(encodingId string, queryParams ...func(*NotificationsWebhooksEncodingEncodingsFinishedAPIListByEncodingIdQueryParams)) (*pagination.WebhooksListByEncodingIdPagination, error) {
	queryParameters := &NotificationsWebhooksEncodingEncodingsFinishedAPIListByEncodingIdQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WebhooksListByEncodingIdPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/finished", nil, &responseModel, reqParams)
	return &responseModel, err
}

// NotificationsWebhooksEncodingEncodingsFinishedAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsWebhooksEncodingEncodingsFinishedAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsFinishedAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}

// NotificationsWebhooksEncodingEncodingsFinishedAPIListByEncodingIdQueryParams contains all query parameters for the ListByEncodingId endpoint
type NotificationsWebhooksEncodingEncodingsFinishedAPIListByEncodingIdQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsFinishedAPIListByEncodingIdQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
