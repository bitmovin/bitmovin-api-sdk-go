package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi) Create(encodingId string, streamId string, burnInSubtitleDvbSub model.BurnInSubtitleDvbSub) (*model.BurnInSubtitleDvbSub, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.BurnInSubtitleDvbSub
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub", &burnInSubtitleDvbSub, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi) Delete(encodingId string, streamId string, subtitleId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["subtitle_id"] = subtitleId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub/{subtitle_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi) Get(encodingId string, streamId string, subtitleId string) (*model.BurnInSubtitleDvbSub, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["subtitle_id"] = subtitleId
    }

    var responseModel *model.BurnInSubtitleDvbSub
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub/{subtitle_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi) List(encodingId string, streamId string, queryParams ...func(*query.BurnInSubtitleDvbSubListQueryParams)) (*pagination.BurnInSubtitleDvbSubsListPagination, error) {
    queryParameters := &query.BurnInSubtitleDvbSubListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.BurnInSubtitleDvbSubsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub", nil, &responseModel, reqParams)
    return responseModel, err
}

