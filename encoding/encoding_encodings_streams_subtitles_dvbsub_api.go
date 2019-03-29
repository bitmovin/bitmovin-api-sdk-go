package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsSubtitlesDvbsubApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsStreamsSubtitlesDvbsubCustomdataApi
}

func NewEncodingEncodingsStreamsSubtitlesDvbsubApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsSubtitlesDvbsubApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsSubtitlesDvbsubApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsStreamsSubtitlesDvbsubCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsSubtitlesDvbsubApi) Create(encodingId string, streamId string, streamDvbSubSubtitle model.StreamDvbSubSubtitle) (*model.StreamDvbSubSubtitle, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }
    payload := model.StreamDvbSubSubtitle(streamDvbSubSubtitle)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/subtitles/dvbsub", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsStreamsSubtitlesDvbsubApi) Delete(encodingId string, streamId string, subtitleId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["subtitle_id"] = subtitleId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/subtitles/dvbsub/{subtitle_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsSubtitlesDvbsubApi) Get(encodingId string, streamId string, subtitleId string) (*model.StreamDvbSubSubtitle, error) {
    var resp *model.StreamDvbSubSubtitle
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["subtitle_id"] = subtitleId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/subtitles/dvbsub/{subtitle_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsSubtitlesDvbsubApi) List(encodingId string, streamId string, queryParams ...func(*query.StreamDvbSubSubtitleListQueryParams)) (*pagination.StreamDvbSubSubtitlesListPagination, error) {
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
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/subtitles/dvbsub", &resp, reqParams)
    return resp, err
}
