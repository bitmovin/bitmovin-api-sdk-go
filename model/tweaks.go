package model

// Tweaks model
type Tweaks struct {
	// Different modes for syncing the start and end of audio input streams with the video inputs. This feature does not work with Dolby Digital (Plus) or Dolby Atmos.
	AudioVideoSyncMode AudioVideoSyncMode `json:"audioVideoSyncMode,omitempty"`
}
