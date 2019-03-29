package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsGcsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsGcsCustomdataApi
}

func NewEncodingInputsGcsApi(configs ...func(*common.ApiClient)) (*EncodingInputsGcsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsGcsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsGcsCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsGcsApi) Create(gcsInput model.GcsInput) (*model.GcsInput, error) {
    payload := model.GcsInput(gcsInput)
    
    err := api.apiClient.Post("/encoding/inputs/gcs", &payload)
    return &payload, err
}
func (api *EncodingInputsGcsApi) Delete(inputId string) (*model.GcsInput, error) {
    var resp *model.GcsInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/gcs/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsGcsApi) List(queryParams ...func(*query.GcsInputListQueryParams)) (*pagination.GcsInputsListPagination, error) {
    queryParameters := &query.GcsInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.GcsInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/gcs", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsGcsApi) Get(inputId string) (*model.GcsInput, error) {
    var resp *model.GcsInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/gcs/{input_id}", &resp, reqParams)
    return resp, err
}
