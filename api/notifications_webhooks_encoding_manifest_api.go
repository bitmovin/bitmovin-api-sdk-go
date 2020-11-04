package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsWebhooksEncodingManifestAPI communicates with '/notifications/webhooks/encoding/manifest/{manifest_id}' endpoints
type NotificationsWebhooksEncodingManifestAPI struct {
	apiClient *apiclient.APIClient

	// Error communicates with '/notifications/webhooks/encoding/manifest/error' endpoints
	Error *NotificationsWebhooksEncodingManifestErrorAPI
	// Finished communicates with '/notifications/webhooks/encoding/manifest/finished' endpoints
	Finished *NotificationsWebhooksEncodingManifestFinishedAPI
}

// NewNotificationsWebhooksEncodingManifestAPI constructor for NotificationsWebhooksEncodingManifestAPI that takes options as argument
func NewNotificationsWebhooksEncodingManifestAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingManifestAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingManifestAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingManifestAPIWithClient constructor for NotificationsWebhooksEncodingManifestAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingManifestAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingManifestAPI {
	a := &NotificationsWebhooksEncodingManifestAPI{apiClient: apiClient}
	a.Error = NewNotificationsWebhooksEncodingManifestErrorAPIWithClient(apiClient)
	a.Finished = NewNotificationsWebhooksEncodingManifestFinishedAPIWithClient(apiClient)

	return a
}

// List Webhook Notifications (Specific Manifest)
func (api *NotificationsWebhooksEncodingManifestAPI) List(manifestId string, queryParams ...func(*NotificationsWebhooksEncodingManifestAPIListQueryParams)) (*pagination.NotificationsListPagination, error) {
	queryParameters := &NotificationsWebhooksEncodingManifestAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.NotificationsListPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/manifest/{manifest_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// NotificationsWebhooksEncodingManifestAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsWebhooksEncodingManifestAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsWebhooksEncodingManifestAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
