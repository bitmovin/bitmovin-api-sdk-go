package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsRedundantRtmpApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsRedundantRtmpApi(configs ...func(*common.ApiClient)) (*EncodingInputsRedundantRtmpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsRedundantRtmpApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsRedundantRtmpApi) List(queryParams ...func(*query.RedundantRtmpInputListQueryParams)) (*pagination.RedundantRtmpInputsListPagination, error) {
    queryParameters := &query.RedundantRtmpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.RedundantRtmpInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/redundant-rtmp", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsRedundantRtmpApi) Get(inputId string) (*model.RedundantRtmpInput, error) {
    var resp *model.RedundantRtmpInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/redundant-rtmp/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsRedundantRtmpApi) Create(redundantRtmpInput model.RedundantRtmpInput) (*model.RedundantRtmpInput, error) {
    payload := model.RedundantRtmpInput(redundantRtmpInput)
    
    err := api.apiClient.Post("/encoding/inputs/redundant-rtmp", &payload)
    return &payload, err
}
func (api *EncodingInputsRedundantRtmpApi) Delete(inputId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/redundant-rtmp/{input_id}", &resp, reqParams)
    return resp, err
}
