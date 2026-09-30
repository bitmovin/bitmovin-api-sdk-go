package model

// StreamMetadata model
type StreamMetadata struct {
	// Language of the media contained in the stream, given as an ISO 639-2 code or a BCP 47 tag, for example 'eng' or 'en-US'. If the value is not set, then no metadata tag is set for the media stream.
	Language *string `json:"language,omitempty"`
	// Display name of the Stream, for example to tell apart multiple audio tracks that share the same language. For CMAF muxings it is written as a 'labl' box (ISO/IEC 14496-12) into the user data of the track. Downstream packagers use it for the Label element in DASH manifests and the NAME attribute of EXT-X-MEDIA tags in HLS playlists. If the value is not set, no label is written. Use 'labelLanguage' to declare which language the label itself is written in.
	Label *string `json:"label,omitempty"`
	// Language the 'label' itself is written in, which is not necessarily the language of the media. A Spanish audio track can carry an English label, for example. For CMAF muxings it is written into the 'labl' box next to the label, and downstream packagers use it for the lang attribute of the Label element in DASH manifests. The value is written as given and downstream packagers expose it unchanged, so set a BCP 47 tag such as 'en' when the manifest should carry that form. The value has to be shaped like a BCP 47 tag: a primary subtag of two or three letters, optionally followed by subtags of up to eight letters or digits, each separated by a hyphen, for example 'en', 'en-US' or 'es-419'. If the value is not set and a label is set, 'language' is used as given, for example 'eng'. Without either, the label language is written as 'und' (undetermined). HLS playlists are not affected, as EXT-X-MEDIA has no equivalent attribute.
	LabelLanguage *string `json:"labelLanguage,omitempty"`
	// Identifier of the switching set the Stream belongs to. For CMAF muxings it is written as a 'kind' box with schemeURI urn:dashif:ingest:switchingset_id (DASH-IF Live Media Ingest) into the user data of the track. Downstream packagers group tracks with the same identifier into one switching set and use it in segment URLs. Only letters, digits, hyphens and underscores are allowed. If the value is not set and a label is set, an identifier is derived from the properties of the Stream, including the label. Without a label, no identifier is written. Set it explicitly when segment URLs have to stay stable across configuration updates.
	SwitchingSetId *string `json:"switchingSetId,omitempty"`
}
