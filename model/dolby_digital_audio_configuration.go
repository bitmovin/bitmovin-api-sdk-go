package model

import (
	"encoding/json"
)

// DolbyDigitalAudioConfiguration model
type DolbyDigitalAudioConfiguration struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Can be freely chosen by the user. (required)
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Target bitrate for the encoded audio in bps (required)
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Audio sampling rate in Hz
	Rate *float64 `json:"rate,omitempty"`
	// BitstreamInfo defines metadata parameters contained in the Dolby Digital audio bitstream
	BitstreamInfo *DolbyDigitalBitstreamInfo `json:"bitstreamInfo,omitempty"`
	// Channel layout of the audio codec configuration. <table> <tr><th colspan=2 align=\"left\"> Available values: </th></tr> <tr><td> NONE </td><td> No channel layout </td></tr> <tr><td> MONO </td><td> Channel layout Mono </td></tr> <tr><td> STEREO </td><td> Channel layout Stereo </td></tr> <tr><td> SURROUND </td><td> Channel layout 3.0 (3 front (left + center + right), no LFE) </td></tr> <tr><td> 3.1 </td><td> Channel layout 3.1 (3 front (left + center + right), LFE) </td></tr> <tr><td> BACK_SURROUND </td><td> Channel layout Surround (2 front (left + right), 1 back center, no LFE) </td></tr> <tr><td> BACK_SURROUND_LFE </td><td> Channel layout Surround with LFE (2 front (left + right), 1 back center, LFE) </td></tr> <tr><td> QUAD </td><td> Channel layout Quad (2 front (left + right), 2 back (left + right), no LFE) </td></tr> <tr><td> QUAD_LFE </td><td> Channel layout Quad with LFE (2 front (left + right), 2 back (left + right), LFE) </td></tr> <tr><td> 4.0 </td><td> Channel layout 4.0 (3 front (left + center + right), 1 back center, no LFE) </td></tr> <tr><td> 4.1 </td><td> Channel layout 4.1 (3 front (left + center + right), 1 back center, LFE) </td></tr> <tr><td> 5.0 </td><td> Channel layout 5.0 (3 front (left + center + right), 2 side (left + right), no LFE) </td></tr> <tr><td> 5.1 </td><td> Channel layout 5.1 (3 front (left + center + right), 2 side (left + right), LFE) </td></tr> </table>
	ChannelLayout DolbyDigitalChannelLayout `json:"channelLayout,omitempty"`
	Downmixing    *DolbyDigitalDownmixing   `json:"downmixing,omitempty"`
	// It provides a framework for signaling new evolution framework applications, such as Intelligent Loudness, in each Dolby codec.
	EvolutionFrameworkControl DolbyDigitalEvolutionFrameworkControl `json:"evolutionFrameworkControl,omitempty"`
	// Settings for loudness control (required)
	LoudnessControl *DolbyDigitalLoudnessControl `json:"loudnessControl,omitempty"`
	Preprocessing   *DolbyDigitalPreprocessing   `json:"preprocessing,omitempty"`
}

func (m DolbyDigitalAudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_DD
}
func (m DolbyDigitalAudioConfiguration) MarshalJSON() ([]byte, error) {
	type M DolbyDigitalAudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DD"

	return json.Marshal(x)
}
