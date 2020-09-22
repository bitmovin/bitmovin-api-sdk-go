package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsTsDrmAesAPI communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes' endpoints
type EncodingEncodingsMuxingsTsDrmAesAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsTsDrmAesCustomdataAPI
}

// NewEncodingEncodingsMuxingsTsDrmAesAPI constructor for EncodingEncodingsMuxingsTsDrmAesAPI that takes options as argument
func NewEncodingEncodingsMuxingsTsDrmAesAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTsDrmAesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsTsDrmAesAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTsDrmAesAPIWithClient constructor for EncodingEncodingsMuxingsTsDrmAesAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTsDrmAesAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTsDrmAesAPI {
	a := &EncodingEncodingsMuxingsTsDrmAesAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsTsDrmAesCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add AES encryption configuration to a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmAesAPI) Create(encodingId string, muxingId string, aesEncryptionDrm model.AesEncryptionDrm) (*model.AesEncryptionDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.AesEncryptionDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes", &aesEncryptionDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete AES encryption configuration from a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmAesAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get AES encryption Details of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmAesAPI) Get(encodingId string, muxingId string, drmId string) (*model.AesEncryptionDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.AesEncryptionDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List AES encryption configurations of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmAesAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsTsDrmAesAPIListQueryParams)) (*pagination.AesEncryptionDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsTsDrmAesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AesEncryptionDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsTsDrmAesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsTsDrmAesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsTsDrmAesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
