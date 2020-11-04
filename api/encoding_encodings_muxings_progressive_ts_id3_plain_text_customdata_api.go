package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text/{id3_tag_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI {
	a := &EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Plain Text ID3 Tag Custom Data of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI) Get(encodingId string, muxingId string, id3TagId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["id3_tag_id"] = id3TagId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text/{id3_tag_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
