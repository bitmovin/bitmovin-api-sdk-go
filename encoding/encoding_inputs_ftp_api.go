package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsFtpApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsFtpCustomdataApi
}

func NewEncodingInputsFtpApi(configs ...func(*common.ApiClient)) (*EncodingInputsFtpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsFtpApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsFtpCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsFtpApi) List(queryParams ...func(*query.FtpInputListQueryParams)) (*pagination.FtpInputsListPagination, error) {
    queryParameters := &query.FtpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.FtpInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/ftp", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsFtpApi) Get(inputId string) (*model.FtpInput, error) {
    var resp *model.FtpInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/ftp/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsFtpApi) Create(ftpInput model.FtpInput) (*model.FtpInput, error) {
    payload := model.FtpInput(ftpInput)
    
    err := api.apiClient.Post("/encoding/inputs/ftp", &payload)
    return &payload, err
}
func (api *EncodingInputsFtpApi) Delete(inputId string) (*model.FtpInput, error) {
    var resp *model.FtpInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/ftp/{input_id}", &resp, reqParams)
    return resp, err
}
