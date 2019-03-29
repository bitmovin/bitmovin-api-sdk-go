package model

type ReuploadSettings struct {
	// Interval in seconds to reupload the DASH manifest (minimum value: 30)
	DashManifestInterval *float64 `json:"dashManifestInterval,omitempty"`
	// Interval in seconds to reupload the HLS master file. This is currently not used, as the master file will always be uploaded when one of the playlist files has changed (minimum value: 30)
	HlsManifestInterval *float64 `json:"hlsManifestInterval,omitempty"`
	// The interval in seconds to reupload the init file for segmented muxings (e.g. fMP4, WebM) (minimum value: 30)
	MuxingInitFileInterval *float64 `json:"muxingInitFileInterval,omitempty"`
}

