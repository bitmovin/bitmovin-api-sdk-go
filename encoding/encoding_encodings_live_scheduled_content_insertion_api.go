package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsLiveScheduledContentInsertionApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsLiveScheduledContentInsertionApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsLiveScheduledContentInsertionApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsLiveScheduledContentInsertionApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsLiveScheduledContentInsertionApi) Create(encodingId string, scheduledContentInsertion model.ScheduledContentInsertion) (*model.ScheduledContentInsertion, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.ScheduledContentInsertion
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/scheduled-content-insertion", &scheduledContentInsertion, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsLiveScheduledContentInsertionApi) Delete(encodingId string, contentInsertionId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["content_insertion_id"] = contentInsertionId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/live/scheduled-content-insertion/{content_insertion_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsLiveScheduledContentInsertionApi) DeleteByEncodingId(encodingId string) (error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/live/scheduled-content-insertion", nil, nil, reqParams)
    return err
}

func (api *EncodingEncodingsLiveScheduledContentInsertionApi) List(encodingId string, queryParams ...func(*query.ScheduledContentInsertionListQueryParams)) (*pagination.ScheduledContentInsertionsListPagination, error) {
    queryParameters := &query.ScheduledContentInsertionListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ScheduledContentInsertionsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/scheduled-content-insertion", nil, &responseModel, reqParams)
    return responseModel, err
}

