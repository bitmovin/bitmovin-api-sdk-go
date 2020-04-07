package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsGcsServiceAccountApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsGcsServiceAccountCustomdataApi
}

func NewEncodingOutputsGcsServiceAccountApi(configs ...func(*common.ApiClient)) (*EncodingOutputsGcsServiceAccountApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsGcsServiceAccountApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsGcsServiceAccountCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsGcsServiceAccountApi) Create(gcsServiceAccountOutput model.GcsServiceAccountOutput) (*model.GcsServiceAccountOutput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.GcsServiceAccountOutput
    err := api.apiClient.Post("/encoding/outputs/gcs-service-account", &gcsServiceAccountOutput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsGcsServiceAccountApi) Delete(outputId string) (*model.GcsServiceAccountOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.GcsServiceAccountOutput
    err := api.apiClient.Delete("/encoding/outputs/gcs-service-account/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsGcsServiceAccountApi) Get(outputId string) (*model.GcsServiceAccountOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.GcsServiceAccountOutput
    err := api.apiClient.Get("/encoding/outputs/gcs-service-account/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsGcsServiceAccountApi) List(queryParams ...func(*query.GcsServiceAccountOutputListQueryParams)) (*pagination.GcsServiceAccountOutputsListPagination, error) {
    queryParameters := &query.GcsServiceAccountOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.GcsServiceAccountOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/gcs-service-account", nil, &responseModel, reqParams)
    return responseModel, err
}

