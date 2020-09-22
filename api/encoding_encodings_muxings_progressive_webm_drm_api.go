package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveWebmDrmAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm' endpoints
type EncodingEncodingsMuxingsProgressiveWebmDrmAPI struct {
	apiClient *apiclient.APIClient

	// Cenc communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/cenc' endpoints
	Cenc *EncodingEncodingsMuxingsProgressiveWebmDrmCencAPI
	// Speke communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/speke' endpoints
	Speke *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPI
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmAPI constructor for EncodingEncodingsMuxingsProgressiveWebmDrmAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWebmDrmAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWebmDrmAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWebmDrmAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWebmDrmAPI {
	a := &EncodingEncodingsMuxingsProgressiveWebmDrmAPI{apiClient: apiClient}
	a.Cenc = NewEncodingEncodingsMuxingsProgressiveWebmDrmCencAPIWithClient(apiClient)
	a.Speke = NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeAPIWithClient(apiClient)

	return a
}

// Get DRM Details of a Progressive WEBM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmAPI) Get(encodingId string, muxingId string, drmId string) (*model.Drm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.Drm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all DRMs of Progressive WEBM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmAPI) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel pagination.DrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm", nil, &responseModel, reqParams)
	return &responseModel, err
}
