package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsPackedAudioDrmAesAPI communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/drm/aes' endpoints
type EncodingEncodingsMuxingsPackedAudioDrmAesAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/drm/aes/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI
}

// NewEncodingEncodingsMuxingsPackedAudioDrmAesAPI constructor for EncodingEncodingsMuxingsPackedAudioDrmAesAPI that takes options as argument
func NewEncodingEncodingsMuxingsPackedAudioDrmAesAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsPackedAudioDrmAesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsPackedAudioDrmAesAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsPackedAudioDrmAesAPIWithClient constructor for EncodingEncodingsMuxingsPackedAudioDrmAesAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsPackedAudioDrmAesAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsPackedAudioDrmAesAPI {
	a := &EncodingEncodingsMuxingsPackedAudioDrmAesAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add AES encryption configuration to the Packed Audio muxing
func (api *EncodingEncodingsMuxingsPackedAudioDrmAesAPI) Create(encodingId string, muxingId string, aesEncryptionDrm model.AesEncryptionDrm) (*model.AesEncryptionDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.AesEncryptionDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/drm/aes", &aesEncryptionDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete AES encryption configuration from a Packed Audio muxing
func (api *EncodingEncodingsMuxingsPackedAudioDrmAesAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get AES encryption Details of a Packed Audio muxing
func (api *EncodingEncodingsMuxingsPackedAudioDrmAesAPI) Get(encodingId string, muxingId string, drmId string) (*model.AesEncryptionDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.AesEncryptionDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List AES encryption configurations of a Packed Audio muxing
func (api *EncodingEncodingsMuxingsPackedAudioDrmAesAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsPackedAudioDrmAesAPIListQueryParams)) (*pagination.AesEncryptionDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsPackedAudioDrmAesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AesEncryptionDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/drm/aes", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsPackedAudioDrmAesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsPackedAudioDrmAesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsPackedAudioDrmAesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
