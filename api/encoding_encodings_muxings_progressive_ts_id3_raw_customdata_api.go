package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw/{id3_tag_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI {
	a := &EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Raw ID3 Tag Custom Data of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataAPI) Get(encodingId string, muxingId string, id3TagId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["id3_tag_id"] = id3TagId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw/{id3_tag_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
