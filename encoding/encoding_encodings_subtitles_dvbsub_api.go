package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsSubtitlesDvbsubApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsSubtitlesDvbsubCustomdataApi
}

func NewEncodingEncodingsSubtitlesDvbsubApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsSubtitlesDvbsubApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsSubtitlesDvbsubApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsSubtitlesDvbsubCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsSubtitlesDvbsubApi) Create(encodingId string, dvbSubtitleSidecarDetails model.DvbSubtitleSidecarDetails) (*model.DvbSubtitleSidecarDetails, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.DvbSubtitleSidecarDetails(dvbSubtitleSidecarDetails)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/subtitles/dvbsub", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsSubtitlesDvbsubApi) List(encodingId string, queryParams ...func(*query.DvbSubtitleSidecarDetailsListQueryParams)) (*pagination.DvbSubtitleSidecarDetailssListPagination, error) {
    queryParameters := &query.DvbSubtitleSidecarDetailsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.DvbSubtitleSidecarDetailssListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/subtitles/dvbsub", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsSubtitlesDvbsubApi) Get(encodingId string, subtitleId string) (*model.DvbSubtitleSidecarDetails, error) {
    var resp *model.DvbSubtitleSidecarDetails
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["subtitle_id"] = subtitleId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/subtitles/dvbsub/{subtitle_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsSubtitlesDvbsubApi) Delete(encodingId string, subtitleId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["subtitle_id"] = subtitleId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/subtitles/dvbsub/{subtitle_id}", &resp, reqParams)
    return resp, err
}
