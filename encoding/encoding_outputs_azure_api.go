package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsAzureApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsAzureCustomdataApi
}

func NewEncodingOutputsAzureApi(configs ...func(*common.ApiClient)) (*EncodingOutputsAzureApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsAzureApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsAzureCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsAzureApi) List(queryParams ...func(*query.AzureOutputListQueryParams)) (*pagination.AzureOutputsListPagination, error) {
    queryParameters := &query.AzureOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.AzureOutputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/outputs/azure", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsAzureApi) Delete(outputId string) (*model.AzureOutput, error) {
    var resp *model.AzureOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Delete("/encoding/outputs/azure/{output_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsAzureApi) Create(azureOutput model.AzureOutput) (*model.AzureOutput, error) {
    payload := model.AzureOutput(azureOutput)
    
    err := api.apiClient.Post("/encoding/outputs/azure", &payload)
    return &payload, err
}
func (api *EncodingOutputsAzureApi) Get(outputId string) (*model.AzureOutput, error) {
    var resp *model.AzureOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Get("/encoding/outputs/azure/{output_id}", &resp, reqParams)
    return resp, err
}
