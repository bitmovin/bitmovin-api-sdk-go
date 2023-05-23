package model

// HistoryStream model
type HistoryStream struct {
	// Stream
	Stream *Stream `json:"stream,omitempty"`
	// List of stream filter configurations
	Filters []StreamFilter `json:"filters,omitempty"`
	// List of Burn-In DVB-SUB subtitles
	BurnInSubtitleDvbSubs []BurnInSubtitleDvbSub `json:"burnInSubtitleDvbSubs,omitempty"`
	// List of burn-in SRT configurations
	BurnInSubtitleSrtSubs []BurnInSubtitleSrt `json:"burnInSubtitleSrtSubs,omitempty"`
	// Nexguard file marker watermarking configuration
	NexGuardFileMarker *NexGuardFileMarker `json:"nexGuardFileMarker,omitempty"`
	// List of caption configurations
	SccCaptions []SccCaption `json:"sccCaptions,omitempty"`
	// List of bif configurations
	Bifs []Bif `json:"bifs,omitempty"`
	// List of Dolby Vision Metadata configurations
	DolbyVisionMetadata []DolbyVisionMetadata `json:"dolbyVisionMetadata,omitempty"`
	// List of Thumbnail configurations
	Thumbnails []Thumbnail `json:"thumbnails,omitempty"`
	// List of Sprite configurations
	Sprites []Sprite `json:"sprites,omitempty"`
	// List of PSNR quality metrics
	PsnrMetrics []PsnrQualityMetric `json:"psnrMetrics,omitempty"`
}
