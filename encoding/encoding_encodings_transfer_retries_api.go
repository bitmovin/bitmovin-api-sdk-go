package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsTransferRetriesApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsTransferRetriesApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsTransferRetriesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsTransferRetriesApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsTransferRetriesApi) Create(encodingId string) (*model.TransferRetry, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.TransferRetry
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/transfer-retries", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsTransferRetriesApi) Get(encodingId string, transferRetryId string) (*model.TransferRetry, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["transfer_retry_id"] = transferRetryId
    }

    var responseModel *model.TransferRetry
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/transfer-retries/{transfer_retry_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsTransferRetriesApi) List(encodingId string, queryParams ...func(*query.TransferRetryListQueryParams)) (*pagination.TransferRetrysListPagination, error) {
    queryParameters := &query.TransferRetryListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.TransferRetrysListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/transfer-retries", nil, &responseModel, reqParams)
    return responseModel, err
}

