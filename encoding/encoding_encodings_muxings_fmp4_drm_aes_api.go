package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsFmp4DrmAesApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsFmp4DrmAesCustomdataApi
}

func NewEncodingEncodingsMuxingsFmp4DrmAesApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4DrmAesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4DrmAesApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsFmp4DrmAesCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4DrmAesApi) Create(encodingId string, muxingId string, aesEncryptionDrm model.AesEncryptionDrm) (*model.AesEncryptionDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.AesEncryptionDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes", &aesEncryptionDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4DrmAesApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4DrmAesApi) Get(encodingId string, muxingId string, drmId string) (*model.AesEncryptionDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.AesEncryptionDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4DrmAesApi) List(encodingId string, muxingId string, queryParams ...func(*query.AesEncryptionDrmListQueryParams)) (*pagination.AesEncryptionDrmsListPagination, error) {
    queryParameters := &query.AesEncryptionDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AesEncryptionDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes", nil, &responseModel, reqParams)
    return responseModel, err
}

