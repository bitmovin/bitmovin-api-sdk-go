package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsWebhooksEncodingEncodingsErrorAPI communicates with '/notifications/webhooks/encoding/encodings/error' endpoints
type NotificationsWebhooksEncodingEncodingsErrorAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/notifications/webhooks/encoding/encodings/error/{webhook_id}/customData' endpoints
    Customdata *NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI

}

// NewNotificationsWebhooksEncodingEncodingsErrorAPI constructor for NotificationsWebhooksEncodingEncodingsErrorAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsErrorAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsErrorAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewNotificationsWebhooksEncodingEncodingsErrorAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsErrorAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsErrorAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsErrorAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsErrorAPI {
    a := &NotificationsWebhooksEncodingEncodingsErrorAPI{apiClient: apiClient}
    a.Customdata = NewNotificationsWebhooksEncodingEncodingsErrorCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add Encoding Error Webhook
func (api *NotificationsWebhooksEncodingEncodingsErrorAPI) Create(webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/error", &webhook, &responseModel, reqParams)
    return &responseModel, err
}
// CreateByEncodingId Add Encoding Error Webhook for specific Encoding Resource. **Note:** A maximum number of 5 webhooks is allowed
func (api *NotificationsWebhooksEncodingEncodingsErrorAPI) CreateByEncodingId(encodingId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/error", &webhook, &responseModel, reqParams)
    return &responseModel, err
}
// DeleteByEncodingIdAndWebhookId Delete Encoding Error Webhook for specific Encoding Resource
func (api *NotificationsWebhooksEncodingEncodingsErrorAPI) DeleteByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/{encoding_id}/error/{webhook_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// DeleteByWebhookId Delete Encoding Error Webhook
func (api *NotificationsWebhooksEncodingEncodingsErrorAPI) DeleteByWebhookId(webhookId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/error/{webhook_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// GetByEncodingIdAndWebhookId Encoding Error Webhook Details for specific Encoding Resource
func (api *NotificationsWebhooksEncodingEncodingsErrorAPI) GetByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.Webhook, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel model.Webhook
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/error/{webhook_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// GetByWebhookId Encoding Error Webhook Details
func (api *NotificationsWebhooksEncodingEncodingsErrorAPI) GetByWebhookId(webhookId string) (*model.Webhook, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel model.Webhook
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/error/{webhook_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Encoding Error Webhooks
func (api *NotificationsWebhooksEncodingEncodingsErrorAPI) List(queryParams ...func(*NotificationsWebhooksEncodingEncodingsErrorAPIListQueryParams)) (*pagination.WebhooksListPagination, error) {
    queryParameters := &NotificationsWebhooksEncodingEncodingsErrorAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.WebhooksListPagination
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/error", nil, &responseModel, reqParams)
    return &responseModel, err
}
// ListByEncodingId List Encoding Error Webhooks for specific Encoding Resource
func (api *NotificationsWebhooksEncodingEncodingsErrorAPI) ListByEncodingId(encodingId string, queryParams ...func(*NotificationsWebhooksEncodingEncodingsErrorAPIListByEncodingIdQueryParams)) (*pagination.WebhooksListByEncodingIdPagination, error) {
    queryParameters := &NotificationsWebhooksEncodingEncodingsErrorAPIListByEncodingIdQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.WebhooksListByEncodingIdPagination
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/error", nil, &responseModel, reqParams)
    return &responseModel, err
}

// NotificationsWebhooksEncodingEncodingsErrorAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsWebhooksEncodingEncodingsErrorAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsErrorAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


// NotificationsWebhooksEncodingEncodingsErrorAPIListByEncodingIdQueryParams contains all query parameters for the ListByEncodingId endpoint
type NotificationsWebhooksEncodingEncodingsErrorAPIListByEncodingIdQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingEncodingsErrorAPIListByEncodingIdQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


