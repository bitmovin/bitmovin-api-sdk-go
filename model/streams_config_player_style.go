package model

// Player style config
type StreamsConfigPlayerStyle struct {
	// Playback marker background color
	PlaybackMarkerBgColor *string `json:"playbackMarkerBgColor,omitempty"`
	// Playback marker border color
	PlaybackMarkerBorderColor *string `json:"playbackMarkerBorderColor,omitempty"`
	// Playback track played color
	PlaybackTrackPlayedColor *string `json:"playbackTrackPlayedColor,omitempty"`
	// Playback track buffered color
	PlaybackTrackBufferedColor *string `json:"playbackTrackBufferedColor,omitempty"`
	// Playback track background color
	PlaybackTrackBgColor *string `json:"playbackTrackBgColor,omitempty"`
	// Text color
	TextColor *string `json:"textColor,omitempty"`
	// Background color
	BackgroundColor *string `json:"backgroundColor,omitempty"`
}
