package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsGcsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsGcsCustomdataApi
}

func NewEncodingOutputsGcsApi(configs ...func(*common.ApiClient)) (*EncodingOutputsGcsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsGcsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsGcsCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsGcsApi) List(queryParams ...func(*query.GcsOutputListQueryParams)) (*pagination.GcsOutputsListPagination, error) {
    queryParameters := &query.GcsOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.GcsOutputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/outputs/gcs", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsGcsApi) Delete(outputId string) (*model.GcsOutput, error) {
    var resp *model.GcsOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Delete("/encoding/outputs/gcs/{output_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingOutputsGcsApi) Create(gcsOutput model.GcsOutput) (*model.GcsOutput, error) {
    payload := model.GcsOutput(gcsOutput)
    
    err := api.apiClient.Post("/encoding/outputs/gcs", &payload)
    return &payload, err
}
func (api *EncodingOutputsGcsApi) Get(outputId string) (*model.GcsOutput, error) {
    var resp *model.GcsOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Get("/encoding/outputs/gcs/{output_id}", &resp, reqParams)
    return resp, err
}
