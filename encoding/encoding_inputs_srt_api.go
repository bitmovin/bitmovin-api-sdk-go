package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsSrtApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsSrtCustomdataApi
}

func NewEncodingInputsSrtApi(configs ...func(*common.ApiClient)) (*EncodingInputsSrtApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsSrtApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsSrtCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsSrtApi) Get(inputId string) (*model.SrtInput, error) {
    var resp *model.SrtInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/srt/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsSrtApi) Delete(inputId string) (*model.SrtInput, error) {
    var resp *model.SrtInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/srt/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsSrtApi) Create(srtInput model.SrtInput) (*model.SrtInput, error) {
    payload := model.SrtInput(srtInput)
    
    err := api.apiClient.Post("/encoding/inputs/srt", &payload)
    return &payload, err
}
func (api *EncodingInputsSrtApi) List(queryParams ...func(*query.SrtInputListQueryParams)) (*pagination.SrtInputsListPagination, error) {
    queryParameters := &query.SrtInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SrtInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/srt", &resp, reqParams)
    return resp, err
}
