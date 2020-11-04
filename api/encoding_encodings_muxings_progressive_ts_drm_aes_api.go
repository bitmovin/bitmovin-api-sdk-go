package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveTsDrmAesAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/aes' endpoints
type EncodingEncodingsMuxingsProgressiveTsDrmAesAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/aes/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsProgressiveTsDrmAesCustomdataAPI
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmAesAPI constructor for EncodingEncodingsMuxingsProgressiveTsDrmAesAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmAesAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsDrmAesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveTsDrmAesAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmAesAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsDrmAesAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmAesAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsDrmAesAPI {
	a := &EncodingEncodingsMuxingsProgressiveTsDrmAesAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsProgressiveTsDrmAesCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add AES encryption configuration to a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmAesAPI) Create(encodingId string, muxingId string, aesEncryptionDrm model.AesEncryptionDrm) (*model.AesEncryptionDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.AesEncryptionDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/aes", &aesEncryptionDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete AES encryption configuration from a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmAesAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get AES encryption Details of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmAesAPI) Get(encodingId string, muxingId string, drmId string) (*model.AesEncryptionDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.AesEncryptionDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List AES encryption configurations of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmAesAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveTsDrmAesAPIListQueryParams)) (*pagination.AesEncryptionDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsProgressiveTsDrmAesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AesEncryptionDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/aes", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveTsDrmAesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveTsDrmAesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveTsDrmAesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
