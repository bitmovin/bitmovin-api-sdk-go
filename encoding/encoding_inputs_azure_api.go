package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsAzureApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsAzureCustomdataApi
}

func NewEncodingInputsAzureApi(configs ...func(*common.ApiClient)) (*EncodingInputsAzureApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsAzureApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsAzureCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsAzureApi) Create(azureInput model.AzureInput) (*model.AzureInput, error) {
    payload := model.AzureInput(azureInput)
    
    err := api.apiClient.Post("/encoding/inputs/azure", &payload)
    return &payload, err
}
func (api *EncodingInputsAzureApi) List(queryParams ...func(*query.AzureInputListQueryParams)) (*pagination.AzureInputsListPagination, error) {
    queryParameters := &query.AzureInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.AzureInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/azure", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsAzureApi) Get(inputId string) (*model.AzureInput, error) {
    var resp *model.AzureInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/azure/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsAzureApi) Delete(inputId string) (*model.AzureInput, error) {
    var resp *model.AzureInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/azure/{input_id}", &resp, reqParams)
    return resp, err
}
