package model

type AnalysisDetails struct {
	VideoStreams []VideoStreamDetails `json:"videoStreams,omitempty"`
	AudioStreams []AudioStreamDetails `json:"audioStreams,omitempty"`
	MetaStreams []MetaStreamDetails `json:"metaStreams,omitempty"`
	SubtitleStreams []SubtitleStreamDetails `json:"subtitleStreams,omitempty"`
}

