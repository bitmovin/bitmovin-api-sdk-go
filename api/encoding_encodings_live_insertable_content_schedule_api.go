package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveInsertableContentScheduleAPI communicates with '/encoding/encodings/{encoding_id}/live/insertable-content/{content_id}/schedule' endpoints
type EncodingEncodingsLiveInsertableContentScheduleAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveInsertableContentScheduleAPI constructor for EncodingEncodingsLiveInsertableContentScheduleAPI that takes options as argument
func NewEncodingEncodingsLiveInsertableContentScheduleAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveInsertableContentScheduleAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsLiveInsertableContentScheduleAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveInsertableContentScheduleAPIWithClient constructor for EncodingEncodingsLiveInsertableContentScheduleAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveInsertableContentScheduleAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveInsertableContentScheduleAPI {
    a := &EncodingEncodingsLiveInsertableContentScheduleAPI{apiClient: apiClient}
    return a
}

// Create Schedule Insertable Content For a Live Encoding
func (api *EncodingEncodingsLiveInsertableContentScheduleAPI) Create(encodingId string, contentId string, scheduledInsertableContent model.ScheduledInsertableContent) (*model.ScheduledInsertableContent, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["content_id"] = contentId
    }

    var responseModel model.ScheduledInsertableContent
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/insertable-content/{content_id}/schedule", &scheduledInsertableContent, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Deschedule Insertable Content
func (api *EncodingEncodingsLiveInsertableContentScheduleAPI) Delete(encodingId string, contentId string, scheduledContentId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["content_id"] = contentId
        params.PathParams["scheduled_content_id"] = scheduledContentId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/live/insertable-content/{content_id}/schedule/{scheduled_content_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}

