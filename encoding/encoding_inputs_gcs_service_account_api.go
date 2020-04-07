package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsGcsServiceAccountApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsGcsServiceAccountCustomdataApi
}

func NewEncodingInputsGcsServiceAccountApi(configs ...func(*common.ApiClient)) (*EncodingInputsGcsServiceAccountApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsGcsServiceAccountApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsGcsServiceAccountCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsGcsServiceAccountApi) Create(gcsServiceAccountInput model.GcsServiceAccountInput) (*model.GcsServiceAccountInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.GcsServiceAccountInput
    err := api.apiClient.Post("/encoding/inputs/gcs-service-account", &gcsServiceAccountInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsGcsServiceAccountApi) Delete(inputId string) (*model.GcsServiceAccountInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.GcsServiceAccountInput
    err := api.apiClient.Delete("/encoding/inputs/gcs-service-account/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsGcsServiceAccountApi) Get(inputId string) (*model.GcsServiceAccountInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.GcsServiceAccountInput
    err := api.apiClient.Get("/encoding/inputs/gcs-service-account/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsGcsServiceAccountApi) List(queryParams ...func(*query.GcsServiceAccountInputListQueryParams)) (*pagination.GcsServiceAccountInputsListPagination, error) {
    queryParameters := &query.GcsServiceAccountInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.GcsServiceAccountInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/gcs-service-account", nil, &responseModel, reqParams)
    return responseModel, err
}

