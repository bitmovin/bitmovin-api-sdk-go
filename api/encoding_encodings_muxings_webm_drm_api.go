package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsWebmDrmAPI communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm' endpoints
type EncodingEncodingsMuxingsWebmDrmAPI struct {
	apiClient *apiclient.APIClient

	// Cenc communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc' endpoints
	Cenc *EncodingEncodingsMuxingsWebmDrmCencAPI
	// Speke communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke' endpoints
	Speke *EncodingEncodingsMuxingsWebmDrmSpekeAPI
}

// NewEncodingEncodingsMuxingsWebmDrmAPI constructor for EncodingEncodingsMuxingsWebmDrmAPI that takes options as argument
func NewEncodingEncodingsMuxingsWebmDrmAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsWebmDrmAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsWebmDrmAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsWebmDrmAPIWithClient constructor for EncodingEncodingsMuxingsWebmDrmAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsWebmDrmAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsWebmDrmAPI {
	a := &EncodingEncodingsMuxingsWebmDrmAPI{apiClient: apiClient}
	a.Cenc = NewEncodingEncodingsMuxingsWebmDrmCencAPIWithClient(apiClient)
	a.Speke = NewEncodingEncodingsMuxingsWebmDrmSpekeAPIWithClient(apiClient)

	return a
}

// Get DRM Details of a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmAPI) Get(encodingId string, muxingId string, drmId string) (*model.Drm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.Drm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all DRM configurations of WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmAPI) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel pagination.DrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm", nil, &responseModel, reqParams)
	return &responseModel, err
}
