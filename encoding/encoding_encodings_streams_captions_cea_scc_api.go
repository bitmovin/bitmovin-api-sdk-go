package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsCaptionsCeaSccApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsStreamsCaptionsCeaSccCustomdataApi
}

func NewEncodingEncodingsStreamsCaptionsCeaSccApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsCaptionsCeaSccApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsCaptionsCeaSccApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsStreamsCaptionsCeaSccCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsCaptionsCeaSccApi) List(encodingId string, streamId string, queryParams ...func(*query.SccCaptionListQueryParams)) (*pagination.SccCaptionsListPagination, error) {
    queryParameters := &query.SccCaptionListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SccCaptionsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsCaptionsCeaSccApi) Create(encodingId string, streamId string, sccCaption model.SccCaption) (*model.SccCaption, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }
    payload := model.SccCaption(sccCaption)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsStreamsCaptionsCeaSccApi) Get(encodingId string, streamId string, captionsId string) (*model.SccCaption, error) {
    var resp *model.SccCaption
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsCaptionsCeaSccApi) Delete(encodingId string, streamId string, captionsId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc/{captions_id}", &resp, reqParams)
    return resp, err
}
