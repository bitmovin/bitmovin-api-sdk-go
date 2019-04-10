package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsCmafDrmFairplayApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsCmafDrmFairplayCustomdataApi
}

func NewEncodingEncodingsMuxingsCmafDrmFairplayApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafDrmFairplayApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafDrmFairplayApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsCmafDrmFairplayCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafDrmFairplayApi) Get(encodingId string, muxingId string, drmId string) (*model.FairPlayDrm, error) {
    var resp *model.FairPlayDrm
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/fairplay/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmFairplayApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/fairplay/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmFairplayApi) Create(encodingId string, muxingId string, fairPlayDrm model.FairPlayDrm) (*model.FairPlayDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.FairPlayDrm(fairPlayDrm)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/fairplay", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsCmafDrmFairplayApi) List(encodingId string, muxingId string, queryParams ...func(*query.FairPlayDrmListQueryParams)) (*pagination.FairPlayDrmsListPagination, error) {
    queryParameters := &query.FairPlayDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.FairPlayDrmsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/fairplay", &resp, reqParams)
    return resp, err
}
