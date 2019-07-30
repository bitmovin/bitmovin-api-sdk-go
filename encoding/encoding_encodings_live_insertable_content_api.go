package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsLiveInsertableContentApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsLiveInsertableContentApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsLiveInsertableContentApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsLiveInsertableContentApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsLiveInsertableContentApi) Create(encodingId string, insertableContent model.InsertableContent) (*model.InsertableContent, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.InsertableContent
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/insertable-content", &insertableContent, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsLiveInsertableContentApi) List(encodingId string, queryParams ...func(*query.InsertableContentListQueryParams)) (*pagination.InsertableContentsListPagination, error) {
    queryParameters := &query.InsertableContentListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.InsertableContentsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/insertable-content", nil, &responseModel, reqParams)
    return responseModel, err
}

