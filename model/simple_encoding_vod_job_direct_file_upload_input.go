package model

import (
	"encoding/json"
)

// SimpleEncodingVodJobDirectFileUploadInput model
type SimpleEncodingVodJobDirectFileUploadInput struct {
	// Id of a DirectFileUploadInput (required)
	InputId *string `json:"inputId,omitempty"`
	// Defines the type of the input file, if no type is set it is assumed that the input file contains at least one video stream and optionally one or multiple audio streams.  Note that when defining video and audio inputs, you can either - add one single input without inputType, in which case that source file must contain a video stream and (if you want audio) one audio stream, or - add one single input with inputType=VIDEO and (if you want audio) one or more inputs with inputType=AUDIO (each containing one audio stream)  Other combinations are not valid.
	InputType SimpleEncodingVodJobInputType `json:"inputType,omitempty"`
	// The language of the audio track, the subtitles, or closed captions file. The language code  must be compliant with [BCP 47](https://datatracker.ietf.org/doc/html/rfc5646).  This property is mandatory if the input provided is of type SUBTITLES or CLOSED_CAPTIONS and  recommended for input type AUDIO and an input that does not specify an input type (combined  audio and video). If an audio input does not specify the language, it is defaulted to `und`  (undefined).
	Language *string `json:"language,omitempty"`
}

func (m SimpleEncodingVodJobDirectFileUploadInput) SimpleEncodingVodJobInputSourceType() SimpleEncodingVodJobInputSourceType {
	return SimpleEncodingVodJobInputSourceType_DIRECT_FILE_UPLOAD
}
func (m SimpleEncodingVodJobDirectFileUploadInput) MarshalJSON() ([]byte, error) {
	type M SimpleEncodingVodJobDirectFileUploadInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DIRECT_FILE_UPLOAD"

	return json.Marshal(x)
}
