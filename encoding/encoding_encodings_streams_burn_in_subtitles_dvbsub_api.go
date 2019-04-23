package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsStreamsBurnInSubtitlesDvbsubCustomdataApi
}

func NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi) List(encodingId string, streamId string, queryParams ...func(*query.StreamDvbSubSubtitleListQueryParams)) (*pagination.StreamDvbSubSubtitlesListPagination, error) {
    queryParameters := &query.StreamDvbSubSubtitleListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.StreamDvbSubSubtitlesListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi) Create(encodingId string, streamId string, streamDvbSubSubtitle model.StreamDvbSubSubtitle) (*model.StreamDvbSubSubtitle, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }
    payload := model.StreamDvbSubSubtitle(streamDvbSubSubtitle)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi) Get(encodingId string, streamId string, subtitleId string) (*model.StreamDvbSubSubtitle, error) {
    var resp *model.StreamDvbSubSubtitle
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["subtitle_id"] = subtitleId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub/{subtitle_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi) Delete(encodingId string, streamId string, subtitleId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["subtitle_id"] = subtitleId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub/{subtitle_id}", &resp, reqParams)
    return resp, err
}
