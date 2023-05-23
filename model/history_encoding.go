package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// HistoryEncoding model
type HistoryEncoding struct {
	// Encoding
	Encoding *Encoding `json:"encoding,omitempty"`
	// Live Details
	Live *LiveEncoding `json:"live,omitempty"`
	// VOD Encoding Start Request
	VodStartReqeust *StartEncodingRequest `json:"vodStartReqeust,omitempty"`
	// Live Encoding Start Request
	LiveStartReqeust *StartLiveEncodingRequest `json:"liveStartReqeust,omitempty"`
	// Encoding Status
	Status             *ModelTask          `json:"status,omitempty"`
	InputStreams       []StreamInput       `json:"inputStreams,omitempty"`
	Streams            []HistoryStream     `json:"streams,omitempty"`
	Muxings            []HistoryMuxing     `json:"muxings,omitempty"`
	KeyFrames          []Keyframe          `json:"keyFrames,omitempty"`
	SidecarFiles       []SidecarFile       `json:"sidecarFiles,omitempty"`
	TransferRetries    []TransferRetry     `json:"transferRetries,omitempty"`
	ConvertSccCaptions []ConvertSccCaption `json:"convertSccCaptions,omitempty"`
}

// UnmarshalJSON unmarshals model HistoryEncoding from a JSON structure
func (m *HistoryEncoding) UnmarshalJSON(raw []byte) error {
	var data struct {
		Encoding           *Encoding                 `json:"encoding"`
		Live               *LiveEncoding             `json:"live"`
		VodStartReqeust    *StartEncodingRequest     `json:"vodStartReqeust"`
		LiveStartReqeust   *StartLiveEncodingRequest `json:"liveStartReqeust"`
		Status             *ModelTask                `json:"status"`
		InputStreams       []StreamInput             `json:"inputStreams"`
		Streams            []HistoryStream           `json:"streams"`
		Muxings            []HistoryMuxing           `json:"muxings"`
		KeyFrames          []Keyframe                `json:"keyFrames"`
		SidecarFiles       json.RawMessage           `json:"sidecarFiles"`
		TransferRetries    []TransferRetry           `json:"transferRetries"`
		ConvertSccCaptions []ConvertSccCaption       `json:"convertSccCaptions"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result HistoryEncoding

	result.Encoding = data.Encoding
	result.Live = data.Live
	result.VodStartReqeust = data.VodStartReqeust
	result.LiveStartReqeust = data.LiveStartReqeust
	result.Status = data.Status
	result.InputStreams = data.InputStreams
	result.Streams = data.Streams
	result.Muxings = data.Muxings
	result.KeyFrames = data.KeyFrames
	result.TransferRetries = data.TransferRetries
	result.ConvertSccCaptions = data.ConvertSccCaptions

	allOfSidecarFiles, err := UnmarshalSidecarFileSlice(bytes.NewBuffer(data.SidecarFiles), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.SidecarFiles = allOfSidecarFiles

	*m = result

	return nil
}
