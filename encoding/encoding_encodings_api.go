package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsApi struct {
    apiClient *common.ApiClient
    Live *EncodingEncodingsLiveApi
    Customdata *EncodingEncodingsCustomdataApi
    Streams *EncodingEncodingsStreamsApi
    InputStreams *EncodingEncodingsInputStreamsApi
    Muxings *EncodingEncodingsMuxingsApi
    Subtitles *EncodingEncodingsSubtitlesApi
    Captions *EncodingEncodingsCaptionsApi
    Sidecars *EncodingEncodingsSidecarsApi
    Keyframes *EncodingEncodingsKeyframesApi
}

func NewEncodingEncodingsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsApi{apiClient: apiClient}

    liveApi, err := NewEncodingEncodingsLiveApi(configs...)
    api.Live = liveApi
    customdataApi, err := NewEncodingEncodingsCustomdataApi(configs...)
    api.Customdata = customdataApi
    streamsApi, err := NewEncodingEncodingsStreamsApi(configs...)
    api.Streams = streamsApi
    inputStreamsApi, err := NewEncodingEncodingsInputStreamsApi(configs...)
    api.InputStreams = inputStreamsApi
    muxingsApi, err := NewEncodingEncodingsMuxingsApi(configs...)
    api.Muxings = muxingsApi
    subtitlesApi, err := NewEncodingEncodingsSubtitlesApi(configs...)
    api.Subtitles = subtitlesApi
    captionsApi, err := NewEncodingEncodingsCaptionsApi(configs...)
    api.Captions = captionsApi
    sidecarsApi, err := NewEncodingEncodingsSidecarsApi(configs...)
    api.Sidecars = sidecarsApi
    keyframesApi, err := NewEncodingEncodingsKeyframesApi(configs...)
    api.Keyframes = keyframesApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsApi) Create(encoding model.Encoding) (*model.Encoding, error) {
    payload := model.Encoding(encoding)
    
    err := api.apiClient.Post("/encoding/encodings", &payload)
    return &payload, err
}
func (api *EncodingEncodingsApi) GetStartRequest(encodingId string) (*model.StartEncodingRequest, error) {
    var resp *model.StartEncodingRequest
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/start", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsApi) Delete(encodingId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsApi) Reprioritize(encodingId string, reprioritizeEncodingRequest model.BitmovinResponse) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.BitmovinResponse(reprioritizeEncodingRequest)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/reprioritize", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsApi) Stop(encodingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    
    var payload model.BitmovinResponse
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/stop", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsApi) Get(encodingId string) (*model.Encoding, error) {
    var resp *model.Encoding
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsApi) Status(encodingId string) (*model.ModelTask, error) {
    var resp *model.ModelTask
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/status", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsApi) Reschedule(encodingId string, rescheduleEncodingRequest model.BitmovinResponse) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.BitmovinResponse(rescheduleEncodingRequest)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/reschedule", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsApi) List(queryParams ...func(*query.EncodingListQueryParams)) (*pagination.EncodingsListPagination, error) {
    queryParameters := &query.EncodingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.EncodingsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsApi) Start(encodingId string, startEncodingRequest model.StartEncodingRequest) (*model.StartEncodingRequest, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.StartEncodingRequest(startEncodingRequest)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/start", &payload, reqParams)
    return &payload, err
}
