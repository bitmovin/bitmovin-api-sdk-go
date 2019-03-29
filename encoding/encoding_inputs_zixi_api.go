package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsZixiApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsZixiCustomdataApi
}

func NewEncodingInputsZixiApi(configs ...func(*common.ApiClient)) (*EncodingInputsZixiApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsZixiApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsZixiCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsZixiApi) Get(inputId string) (*model.ZixiInput, error) {
    var resp *model.ZixiInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/zixi/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsZixiApi) List(queryParams ...func(*query.ZixiInputListQueryParams)) (*pagination.ZixiInputsListPagination, error) {
    queryParameters := &query.ZixiInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.ZixiInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/zixi", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsZixiApi) Delete(inputId string) (*model.ZixiInput, error) {
    var resp *model.ZixiInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/zixi/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsZixiApi) Create(zixiInput model.ZixiInput) (*model.ZixiInput, error) {
    payload := model.ZixiInput(zixiInput)
    
    err := api.apiClient.Post("/encoding/inputs/zixi", &payload)
    return &payload, err
}
