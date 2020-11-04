package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingConfigurationsAudioAPI intermediary API object with no endpoints
type EncodingConfigurationsAudioAPI struct {
	apiClient *apiclient.APIClient

	// Aac communicates with '/encoding/configurations/audio/aac' endpoints
	Aac *EncodingConfigurationsAudioAacAPI
	// DtsPassthrough communicates with '/encoding/configurations/audio/dts-passthrough' endpoints
	DtsPassthrough *EncodingConfigurationsAudioDtsPassthroughAPI
	// DolbyAtmos communicates with '/encoding/configurations/audio/dolby-atmos' endpoints
	DolbyAtmos *EncodingConfigurationsAudioDolbyAtmosAPI
	// HeAacV1 communicates with '/encoding/configurations/audio/he-aac-v1' endpoints
	HeAacV1 *EncodingConfigurationsAudioHeAacV1API
	// HeAacV2 communicates with '/encoding/configurations/audio/he-aac-v2' endpoints
	HeAacV2 *EncodingConfigurationsAudioHeAacV2API
	// Vorbis communicates with '/encoding/configurations/audio/vorbis' endpoints
	Vorbis *EncodingConfigurationsAudioVorbisAPI
	// Opus communicates with '/encoding/configurations/audio/opus' endpoints
	Opus *EncodingConfigurationsAudioOpusAPI
	// Pcm communicates with '/encoding/configurations/audio/pcm' endpoints
	Pcm *EncodingConfigurationsAudioPcmAPI
	// Ac3 communicates with '/encoding/configurations/audio/ac3' endpoints
	Ac3 *EncodingConfigurationsAudioAc3API
	// Eac3 communicates with '/encoding/configurations/audio/eac3' endpoints
	Eac3 *EncodingConfigurationsAudioEac3API
	// Mp2 communicates with '/encoding/configurations/audio/mp2' endpoints
	Mp2 *EncodingConfigurationsAudioMp2API
	// Mp3 communicates with '/encoding/configurations/audio/mp3' endpoints
	Mp3 *EncodingConfigurationsAudioMp3API
}

// NewEncodingConfigurationsAudioAPI constructor for EncodingConfigurationsAudioAPI that takes options as argument
func NewEncodingConfigurationsAudioAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioAPIWithClient constructor for EncodingConfigurationsAudioAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioAPI {
	a := &EncodingConfigurationsAudioAPI{apiClient: apiClient}
	a.Aac = NewEncodingConfigurationsAudioAacAPIWithClient(apiClient)
	a.DtsPassthrough = NewEncodingConfigurationsAudioDtsPassthroughAPIWithClient(apiClient)
	a.DolbyAtmos = NewEncodingConfigurationsAudioDolbyAtmosAPIWithClient(apiClient)
	a.HeAacV1 = NewEncodingConfigurationsAudioHeAacV1APIWithClient(apiClient)
	a.HeAacV2 = NewEncodingConfigurationsAudioHeAacV2APIWithClient(apiClient)
	a.Vorbis = NewEncodingConfigurationsAudioVorbisAPIWithClient(apiClient)
	a.Opus = NewEncodingConfigurationsAudioOpusAPIWithClient(apiClient)
	a.Pcm = NewEncodingConfigurationsAudioPcmAPIWithClient(apiClient)
	a.Ac3 = NewEncodingConfigurationsAudioAc3APIWithClient(apiClient)
	a.Eac3 = NewEncodingConfigurationsAudioEac3APIWithClient(apiClient)
	a.Mp2 = NewEncodingConfigurationsAudioMp2APIWithClient(apiClient)
	a.Mp3 = NewEncodingConfigurationsAudioMp3APIWithClient(apiClient)

	return a
}
