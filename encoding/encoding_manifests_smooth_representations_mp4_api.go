package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsSmoothRepresentationsMp4Api struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsSmoothRepresentationsMp4Api(configs ...func(*common.ApiClient)) (*EncodingManifestsSmoothRepresentationsMp4Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsSmoothRepresentationsMp4Api{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsSmoothRepresentationsMp4Api) Delete(manifestId string, representationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["representation_id"] = representationId
	}
    err := api.apiClient.Delete("/encoding/manifests/smooth/{manifest_id}/representations/mp4/{representation_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsSmoothRepresentationsMp4Api) List(manifestId string, queryParams ...func(*query.SmoothStreamingRepresentationListQueryParams)) (*pagination.SmoothStreamingRepresentationsListPagination, error) {
    queryParameters := &query.SmoothStreamingRepresentationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SmoothStreamingRepresentationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/representations/mp4", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsSmoothRepresentationsMp4Api) Get(manifestId string, representationId string) (*model.SmoothStreamingRepresentation, error) {
    var resp *model.SmoothStreamingRepresentation
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["representation_id"] = representationId
	}
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/representations/mp4/{representation_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsSmoothRepresentationsMp4Api) Create(manifestId string, smoothStreamingRepresentation model.SmoothStreamingRepresentation) (*model.SmoothStreamingRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    payload := model.SmoothStreamingRepresentation(smoothStreamingRepresentation)
    
    err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/representations/mp4", &payload, reqParams)
    return &payload, err
}
