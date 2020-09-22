package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsEmailsEncodingEncodingsAPI communicates with '/notifications/emails/encoding/encodings/{encoding_id}' endpoints
type NotificationsEmailsEncodingEncodingsAPI struct {
	apiClient *apiclient.APIClient

	// LiveInputStreamChanged communicates with '/notifications/emails/encoding/encodings/live-input-stream-changed' endpoints
	LiveInputStreamChanged *NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI
	// Error communicates with '/notifications/emails/encoding/encodings/error' endpoints
	Error *NotificationsEmailsEncodingEncodingsErrorAPI
}

// NewNotificationsEmailsEncodingEncodingsAPI constructor for NotificationsEmailsEncodingEncodingsAPI that takes options as argument
func NewNotificationsEmailsEncodingEncodingsAPI(options ...apiclient.APIClientOption) (*NotificationsEmailsEncodingEncodingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsEmailsEncodingEncodingsAPIWithClient(apiClient), nil
}

// NewNotificationsEmailsEncodingEncodingsAPIWithClient constructor for NotificationsEmailsEncodingEncodingsAPI that takes an APIClient as argument
func NewNotificationsEmailsEncodingEncodingsAPIWithClient(apiClient *apiclient.APIClient) *NotificationsEmailsEncodingEncodingsAPI {
	a := &NotificationsEmailsEncodingEncodingsAPI{apiClient: apiClient}
	a.LiveInputStreamChanged = NewNotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPIWithClient(apiClient)
	a.Error = NewNotificationsEmailsEncodingEncodingsErrorAPIWithClient(apiClient)

	return a
}

// List Email Notifications (Specific Encoding)
func (api *NotificationsEmailsEncodingEncodingsAPI) List(encodingId string, queryParams ...func(*NotificationsEmailsEncodingEncodingsAPIListQueryParams)) (*pagination.EmailNotificationWithStreamConditionssListPagination, error) {
	queryParameters := &NotificationsEmailsEncodingEncodingsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EmailNotificationWithStreamConditionssListPagination
	err := api.apiClient.Get("/notifications/emails/encoding/encodings/{encoding_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// NotificationsEmailsEncodingEncodingsAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsEmailsEncodingEncodingsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsEmailsEncodingEncodingsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
