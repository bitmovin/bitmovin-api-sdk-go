package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsUdpMulticastApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsUdpMulticastCustomdataApi
}

func NewEncodingInputsUdpMulticastApi(configs ...func(*common.ApiClient)) (*EncodingInputsUdpMulticastApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsUdpMulticastApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsUdpMulticastCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsUdpMulticastApi) Create(udpMulticastInput model.UdpMulticastInput) (*model.UdpMulticastInput, error) {
    payload := model.UdpMulticastInput(udpMulticastInput)
    
    err := api.apiClient.Post("/encoding/inputs/udp-multicast", &payload)
    return &payload, err
}
func (api *EncodingInputsUdpMulticastApi) Get(inputId string) (*model.UdpMulticastInput, error) {
    var resp *model.UdpMulticastInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/udp-multicast/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsUdpMulticastApi) Delete(inputId string) (*model.UdpMulticastInput, error) {
    var resp *model.UdpMulticastInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/udp-multicast/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsUdpMulticastApi) List(queryParams ...func(*query.UdpMulticastInputListQueryParams)) (*pagination.UdpMulticastInputsListPagination, error) {
    queryParameters := &query.UdpMulticastInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.UdpMulticastInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/udp-multicast", &resp, reqParams)
    return resp, err
}
