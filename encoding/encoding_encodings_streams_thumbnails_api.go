package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsThumbnailsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsStreamsThumbnailsCustomdataApi
}

func NewEncodingEncodingsStreamsThumbnailsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsThumbnailsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsThumbnailsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsStreamsThumbnailsCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsThumbnailsApi) Create(encodingId string, streamId string, thumbnail model.Thumbnail) (*model.Thumbnail, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }
    payload := model.Thumbnail(thumbnail)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsStreamsThumbnailsApi) Get(encodingId string, streamId string, thumbnailId string) (*model.Thumbnail, error) {
    var resp *model.Thumbnail
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["thumbnail_id"] = thumbnailId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsThumbnailsApi) List(encodingId string, streamId string, queryParams ...func(*query.ThumbnailListQueryParams)) (*pagination.ThumbnailsListPagination, error) {
    queryParameters := &query.ThumbnailListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.ThumbnailsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsThumbnailsApi) Delete(encodingId string, streamId string, thumbnailId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["thumbnail_id"] = thumbnailId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}", &resp, reqParams)
    return resp, err
}
