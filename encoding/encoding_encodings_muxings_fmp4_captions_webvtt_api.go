package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsFmp4CaptionsWebvttApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsFmp4CaptionsWebvttCustomdataApi
}

func NewEncodingEncodingsMuxingsFmp4CaptionsWebvttApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4CaptionsWebvttApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4CaptionsWebvttApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsFmp4CaptionsWebvttCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4CaptionsWebvttApi) Get(encodingId string, muxingId string, captionsId string) (*model.WebVttEmbed, error) {
    var resp *model.WebVttEmbed
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/captions/webvtt/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsFmp4CaptionsWebvttApi) List(encodingId string, muxingId string) (*pagination.WebVttEmbedsListPagination, error) {
    var resp *pagination.WebVttEmbedsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/captions/webvtt", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsFmp4CaptionsWebvttApi) Create(encodingId string, muxingId string, webVttEmbed model.WebVttEmbed) (*model.WebVttEmbed, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.WebVttEmbed(webVttEmbed)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/captions/webvtt", &payload, reqParams)
    return &payload, err
}
