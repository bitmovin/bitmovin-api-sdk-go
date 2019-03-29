package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsSftpApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsSftpCustomdataApi
}

func NewEncodingOutputsSftpApi(configs ...func(*common.ApiClient)) (*EncodingOutputsSftpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsSftpApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsSftpCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsSftpApi) Create(sftpOutput model.SftpOutput) (*model.SftpOutput, error) {
    payload := model.SftpOutput(sftpOutput)
    
    err := api.apiClient.Post("/encoding/outputs/sftp", &payload)
    return &payload, err
}
func (api *EncodingOutputsSftpApi) Get(outputId string) (*model.SftpOutput, error) {
    var resp *model.SftpOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Get("/encoding/outputs/sftp/{output_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsSftpApi) List(queryParams ...func(*query.SftpOutputListQueryParams)) (*pagination.SftpOutputsListPagination, error) {
    queryParameters := &query.SftpOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SftpOutputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/outputs/sftp", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsSftpApi) Delete(outputId string) (*model.SftpOutput, error) {
    var resp *model.SftpOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Delete("/encoding/outputs/sftp/{output_id}", &resp, reqParams)
    return resp, err
}
