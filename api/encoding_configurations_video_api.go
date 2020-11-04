package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingConfigurationsVideoAPI intermediary API object with no endpoints
type EncodingConfigurationsVideoAPI struct {
	apiClient *apiclient.APIClient

	// H262 communicates with '/encoding/configurations/video/h262' endpoints
	H262 *EncodingConfigurationsVideoH262API
	// H264 communicates with '/encoding/configurations/video/h264' endpoints
	H264 *EncodingConfigurationsVideoH264API
	// H265 communicates with '/encoding/configurations/video/h265' endpoints
	H265 *EncodingConfigurationsVideoH265API
	// Vp8 communicates with '/encoding/configurations/video/vp8' endpoints
	Vp8 *EncodingConfigurationsVideoVp8API
	// Vp9 communicates with '/encoding/configurations/video/vp9' endpoints
	Vp9 *EncodingConfigurationsVideoVp9API
	// Av1 communicates with '/encoding/configurations/video/av1' endpoints
	Av1 *EncodingConfigurationsVideoAv1API
	// Mjpeg communicates with '/encoding/configurations/video/mjpeg' endpoints
	Mjpeg *EncodingConfigurationsVideoMjpegAPI
}

// NewEncodingConfigurationsVideoAPI constructor for EncodingConfigurationsVideoAPI that takes options as argument
func NewEncodingConfigurationsVideoAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsVideoAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoAPIWithClient constructor for EncodingConfigurationsVideoAPI that takes an APIClient as argument
func NewEncodingConfigurationsVideoAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoAPI {
	a := &EncodingConfigurationsVideoAPI{apiClient: apiClient}
	a.H262 = NewEncodingConfigurationsVideoH262APIWithClient(apiClient)
	a.H264 = NewEncodingConfigurationsVideoH264APIWithClient(apiClient)
	a.H265 = NewEncodingConfigurationsVideoH265APIWithClient(apiClient)
	a.Vp8 = NewEncodingConfigurationsVideoVp8APIWithClient(apiClient)
	a.Vp9 = NewEncodingConfigurationsVideoVp9APIWithClient(apiClient)
	a.Av1 = NewEncodingConfigurationsVideoAv1APIWithClient(apiClient)
	a.Mjpeg = NewEncodingConfigurationsVideoMjpegAPIWithClient(apiClient)

	return a
}
