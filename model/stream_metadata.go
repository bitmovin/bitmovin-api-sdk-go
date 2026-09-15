package model

// StreamMetadata model
type StreamMetadata struct {
	// Language of the media contained in the stream. If the value is not set, then no metadata tag is set for the media stream.
	Language *string `json:"language,omitempty"`
	// Display name of the Stream, for example to tell apart multiple audio tracks that share the same language. For CMAF muxings it is written as a 'labl' box (ISO/IEC 14496-12) into the user data of the track. Downstream packagers use it for the Label element in DASH manifests and the NAME attribute of EXT-X-MEDIA tags in HLS playlists. If the value is not set, no label is written.
	Label *string `json:"label,omitempty"`
	// Identifier of the switching set the Stream belongs to. For CMAF muxings it is written as a 'kind' box with schemeURI urn:dashif:ingest:switchingset_id (DASH-IF Live Media Ingest) into the user data of the track. Downstream packagers group tracks with the same identifier into one switching set and use it in segment URLs. Only letters, digits, hyphens and underscores are allowed. If the value is not set and a label is set, an identifier is derived from the properties of the Stream, including the label. Without a label, no identifier is written. Set it explicitly when segment URLs have to stay stable across configuration updates.
	SwitchingSetId *string `json:"switchingSetId,omitempty"`
}
