package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsSftpApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsSftpCustomdataApi
}

func NewEncodingInputsSftpApi(configs ...func(*common.ApiClient)) (*EncodingInputsSftpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsSftpApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsSftpCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsSftpApi) Create(sftpInput model.SftpInput) (*model.SftpInput, error) {
    payload := model.SftpInput(sftpInput)
    
    err := api.apiClient.Post("/encoding/inputs/sftp", &payload)
    return &payload, err
}
func (api *EncodingInputsSftpApi) Get(inputId string) (*model.SftpInput, error) {
    var resp *model.SftpInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/sftp/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsSftpApi) List(queryParams ...func(*query.SftpInputListQueryParams)) (*pagination.SftpInputsListPagination, error) {
    queryParameters := &query.SftpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SftpInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/sftp", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsSftpApi) Delete(inputId string) (*model.SftpInput, error) {
    var resp *model.SftpInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/sftp/{input_id}", &resp, reqParams)
    return resp, err
}
