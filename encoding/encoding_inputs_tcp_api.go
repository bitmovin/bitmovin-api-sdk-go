package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsTcpApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsTcpApi(configs ...func(*common.ApiClient)) (*EncodingInputsTcpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsTcpApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsTcpApi) Get(inputId string) (*model.TcpInput, error) {
    var resp *model.TcpInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/tcp/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsTcpApi) List(queryParams ...func(*query.TcpInputListQueryParams)) (*pagination.TcpInputsListPagination, error) {
    queryParameters := &query.TcpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.TcpInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/tcp", &resp, reqParams)
    return resp, err
}
