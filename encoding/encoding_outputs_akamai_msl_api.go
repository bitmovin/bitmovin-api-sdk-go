package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsAkamaiMslApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsAkamaiMslCustomdataApi
}

func NewEncodingOutputsAkamaiMslApi(configs ...func(*common.ApiClient)) (*EncodingOutputsAkamaiMslApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsAkamaiMslApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsAkamaiMslCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsAkamaiMslApi) List(queryParams ...func(*query.AkamaiMslOutputListQueryParams)) (*pagination.AkamaiMslOutputsListPagination, error) {
    queryParameters := &query.AkamaiMslOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.AkamaiMslOutputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/outputs/akamai-msl", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsAkamaiMslApi) Create(akamaiMslOutput model.AkamaiMslOutput) (*model.AkamaiMslOutput, error) {
    payload := model.AkamaiMslOutput(akamaiMslOutput)
    
    err := api.apiClient.Post("/encoding/outputs/akamai-msl", &payload)
    return &payload, err
}
func (api *EncodingOutputsAkamaiMslApi) Get(outputId string) (*model.AkamaiMslOutput, error) {
    var resp *model.AkamaiMslOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Get("/encoding/outputs/akamai-msl/{output_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsAkamaiMslApi) Delete(outputId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Delete("/encoding/outputs/akamai-msl/{output_id}", &resp, reqParams)
    return resp, err
}
