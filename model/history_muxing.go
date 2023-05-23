package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// HistoryMuxing model
type HistoryMuxing struct {
	// Muxing
	Muxing                           *Muxing                           `json:"muxing,omitempty"`
	Drms                             []Drm                             `json:"drms,omitempty"`
	BroadcastTsMuxingInformation     *BroadcastTsMuxingInformation     `json:"broadcastTsMuxingInformation,omitempty"`
	Fmp4MuxingInformation            *Fmp4MuxingInformation            `json:"fmp4MuxingInformation,omitempty"`
	Mp3MuxingInformation             *Mp3MuxingInformation             `json:"mp3MuxingInformation,omitempty"`
	Mp4MuxingInformation             *Mp4MuxingInformation             `json:"mp4MuxingInformation,omitempty"`
	PackedAudioMuxingInformation     *PackedAudioMuxingInformation     `json:"packedAudioMuxingInformation,omitempty"`
	ProgressiveMovMuxingInformation  *ProgressiveMovMuxingInformation  `json:"progressiveMovMuxingInformation,omitempty"`
	ProgressiveTsMuxingInformation   *ProgressiveTsMuxingInformation   `json:"progressiveTsMuxingInformation,omitempty"`
	ProgressiveWebmMuxingInformation *ProgressiveWebmMuxingInformation `json:"progressiveWebmMuxingInformation,omitempty"`
}

// UnmarshalJSON unmarshals model HistoryMuxing from a JSON structure
func (m *HistoryMuxing) UnmarshalJSON(raw []byte) error {
	var data struct {
		Muxing                           json.RawMessage                   `json:"muxing"`
		Drms                             json.RawMessage                   `json:"drms"`
		BroadcastTsMuxingInformation     *BroadcastTsMuxingInformation     `json:"broadcastTsMuxingInformation"`
		Fmp4MuxingInformation            *Fmp4MuxingInformation            `json:"fmp4MuxingInformation"`
		Mp3MuxingInformation             *Mp3MuxingInformation             `json:"mp3MuxingInformation"`
		Mp4MuxingInformation             *Mp4MuxingInformation             `json:"mp4MuxingInformation"`
		PackedAudioMuxingInformation     *PackedAudioMuxingInformation     `json:"packedAudioMuxingInformation"`
		ProgressiveMovMuxingInformation  *ProgressiveMovMuxingInformation  `json:"progressiveMovMuxingInformation"`
		ProgressiveTsMuxingInformation   *ProgressiveTsMuxingInformation   `json:"progressiveTsMuxingInformation"`
		ProgressiveWebmMuxingInformation *ProgressiveWebmMuxingInformation `json:"progressiveWebmMuxingInformation"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result HistoryMuxing

	result.BroadcastTsMuxingInformation = data.BroadcastTsMuxingInformation
	result.Fmp4MuxingInformation = data.Fmp4MuxingInformation
	result.Mp3MuxingInformation = data.Mp3MuxingInformation
	result.Mp4MuxingInformation = data.Mp4MuxingInformation
	result.PackedAudioMuxingInformation = data.PackedAudioMuxingInformation
	result.ProgressiveMovMuxingInformation = data.ProgressiveMovMuxingInformation
	result.ProgressiveTsMuxingInformation = data.ProgressiveTsMuxingInformation
	result.ProgressiveWebmMuxingInformation = data.ProgressiveWebmMuxingInformation

	allOfMuxing, err := UnmarshalMuxing(bytes.NewBuffer(data.Muxing), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Muxing = &allOfMuxing
	allOfDrms, err := UnmarshalDrmSlice(bytes.NewBuffer(data.Drms), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Drms = allOfDrms

	*m = result

	return nil
}
