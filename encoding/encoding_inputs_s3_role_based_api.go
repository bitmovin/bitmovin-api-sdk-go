package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsS3RoleBasedApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsS3RoleBasedCustomdataApi
}

func NewEncodingInputsS3RoleBasedApi(configs ...func(*common.ApiClient)) (*EncodingInputsS3RoleBasedApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsS3RoleBasedApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsS3RoleBasedCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsS3RoleBasedApi) Create(s3RoleBasedInput model.S3RoleBasedInput) (*model.S3RoleBasedInput, error) {
    payload := model.S3RoleBasedInput(s3RoleBasedInput)
    
    err := api.apiClient.Post("/encoding/inputs/s3-role-based", &payload)
    return &payload, err
}
func (api *EncodingInputsS3RoleBasedApi) Delete(inputId string) (*model.S3RoleBasedInput, error) {
    var resp *model.S3RoleBasedInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/s3-role-based/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsS3RoleBasedApi) Get(inputId string) (*model.S3RoleBasedInput, error) {
    var resp *model.S3RoleBasedInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/s3-role-based/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsS3RoleBasedApi) List(queryParams ...func(*query.S3RoleBasedInputListQueryParams)) (*pagination.S3RoleBasedInputsListPagination, error) {
    queryParameters := &query.S3RoleBasedInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.S3RoleBasedInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/s3-role-based", &resp, reqParams)
    return resp, err
}
