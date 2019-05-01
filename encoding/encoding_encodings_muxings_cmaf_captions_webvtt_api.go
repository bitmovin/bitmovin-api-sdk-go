package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsCmafCaptionsWebvttApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsCmafCaptionsWebvttCustomdataApi
}

func NewEncodingEncodingsMuxingsCmafCaptionsWebvttApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafCaptionsWebvttApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafCaptionsWebvttApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsCmafCaptionsWebvttCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafCaptionsWebvttApi) List(encodingId string, muxingId string) (*pagination.WebVttEmbedsListPagination, error) {
    var resp *pagination.WebVttEmbedsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/webvtt", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafCaptionsWebvttApi) Create(encodingId string, muxingId string, webVttEmbed model.WebVttEmbed) (*model.WebVttEmbed, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.WebVttEmbed(webVttEmbed)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/webvtt", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsCmafCaptionsWebvttApi) Get(encodingId string, muxingId string, captionsId string) (*model.WebVttEmbed, error) {
    var resp *model.WebVttEmbed
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/webvtt/{captions_id}", &resp, reqParams)
    return resp, err
}