package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsTsDrmAPI communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm' endpoints
type EncodingEncodingsMuxingsTsDrmAPI struct {
	apiClient *apiclient.APIClient

	// Fairplay communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/fairplay' endpoints
	Fairplay *EncodingEncodingsMuxingsTsDrmFairplayAPI
	// Aes communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes' endpoints
	Aes *EncodingEncodingsMuxingsTsDrmAesAPI
	// Speke communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke' endpoints
	Speke *EncodingEncodingsMuxingsTsDrmSpekeAPI
}

// NewEncodingEncodingsMuxingsTsDrmAPI constructor for EncodingEncodingsMuxingsTsDrmAPI that takes options as argument
func NewEncodingEncodingsMuxingsTsDrmAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTsDrmAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsTsDrmAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTsDrmAPIWithClient constructor for EncodingEncodingsMuxingsTsDrmAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTsDrmAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTsDrmAPI {
	a := &EncodingEncodingsMuxingsTsDrmAPI{apiClient: apiClient}
	a.Fairplay = NewEncodingEncodingsMuxingsTsDrmFairplayAPIWithClient(apiClient)
	a.Aes = NewEncodingEncodingsMuxingsTsDrmAesAPIWithClient(apiClient)
	a.Speke = NewEncodingEncodingsMuxingsTsDrmSpekeAPIWithClient(apiClient)

	return a
}

// Get DRM Details of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmAPI) Get(encodingId string, muxingId string, drmId string) (*model.Drm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.Drm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all DRM configurations of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmAPI) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel pagination.DrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm", nil, &responseModel, reqParams)
	return &responseModel, err
}
