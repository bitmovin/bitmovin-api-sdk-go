package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsHlsMediaTypeAPI communicates with '/encoding/manifests/hls/{manifest_id}/media/{media_id}/type' endpoints
type EncodingManifestsHlsMediaTypeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsMediaTypeAPI constructor for EncodingManifestsHlsMediaTypeAPI that takes options as argument
func NewEncodingManifestsHlsMediaTypeAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsMediaTypeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsHlsMediaTypeAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsMediaTypeAPIWithClient constructor for EncodingManifestsHlsMediaTypeAPI that takes an APIClient as argument
func NewEncodingManifestsHlsMediaTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsMediaTypeAPI {
	a := &EncodingManifestsHlsMediaTypeAPI{apiClient: apiClient}
	return a
}

// Get HLS Media Type
func (api *EncodingManifestsHlsMediaTypeAPI) Get(manifestId string, mediaId string) (*model.MediaInfoTypeResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["media_id"] = mediaId
	}

	var responseModel model.MediaInfoTypeResponse
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/{media_id}/type", nil, &responseModel, reqParams)
	return &responseModel, err
}
