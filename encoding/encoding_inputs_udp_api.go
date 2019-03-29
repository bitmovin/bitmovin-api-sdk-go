package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsUdpApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsUdpApi(configs ...func(*common.ApiClient)) (*EncodingInputsUdpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsUdpApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsUdpApi) List(queryParams ...func(*query.UdpInputListQueryParams)) (*pagination.UdpInputsListPagination, error) {
    queryParameters := &query.UdpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.UdpInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/udp", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsUdpApi) Get(inputId string) (*model.UdpInput, error) {
    var resp *model.UdpInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/udp/{input_id}", &resp, reqParams)
    return resp, err
}
