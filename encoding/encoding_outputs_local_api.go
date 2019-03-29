package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsLocalApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsLocalCustomdataApi
}

func NewEncodingOutputsLocalApi(configs ...func(*common.ApiClient)) (*EncodingOutputsLocalApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsLocalApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsLocalCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsLocalApi) Create(localOutput model.LocalOutput) (*model.LocalOutput, error) {
    payload := model.LocalOutput(localOutput)
    
    err := api.apiClient.Post("/encoding/outputs/local", &payload)
    return &payload, err
}
func (api *EncodingOutputsLocalApi) Get(outputId string) (*model.LocalOutput, error) {
    var resp *model.LocalOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Get("/encoding/outputs/local/{output_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsLocalApi) Delete(outputId string) (*model.LocalOutput, error) {
    var resp *model.LocalOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Delete("/encoding/outputs/local/{output_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsLocalApi) List(queryParams ...func(*query.LocalOutputListQueryParams)) (*pagination.LocalOutputsListPagination, error) {
    queryParameters := &query.LocalOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.LocalOutputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/outputs/local", &resp, reqParams)
    return resp, err
}
