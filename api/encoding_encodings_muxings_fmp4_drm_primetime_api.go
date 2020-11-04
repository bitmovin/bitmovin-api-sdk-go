package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/primetime' endpoints
type EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/primetime/{drm_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI
}

// NewEncodingEncodingsMuxingsFmp4DrmPrimetimeAPI constructor for EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmPrimetimeAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4DrmPrimetimeAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmPrimetimeAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmPrimetimeAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI {
	a := &EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add PrimeTime DRM to an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI) Create(encodingId string, muxingId string, primeTimeDrm model.PrimeTimeDrm) (*model.PrimeTimeDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.PrimeTimeDrm
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/primetime", &primeTimeDrm, &responseModel, reqParams)
	return &responseModel, err
}

// Delete PrimeTime DRM from an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/primetime/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get PrimeTime DRM Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI) Get(encodingId string, muxingId string, drmId string) (*model.PrimeTimeDrm, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.PrimeTimeDrm
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/primetime/{drm_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List PrimeTime DRMs of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4DrmPrimetimeAPIListQueryParams)) (*pagination.PrimeTimeDrmsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsFmp4DrmPrimetimeAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.PrimeTimeDrmsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/primetime", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4DrmPrimetimeAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4DrmPrimetimeAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4DrmPrimetimeAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
