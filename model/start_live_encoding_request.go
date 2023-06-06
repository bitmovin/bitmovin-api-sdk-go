package model

// StartLiveEncodingRequest model
type StartLiveEncodingRequest struct {
	// Key for the stream. (a-zA-Z, 3-20 characters) (required)
	StreamKey *string `json:"streamKey,omitempty"`
	// List of Hls manifests to use for this live encoding
	HlsManifests []LiveHlsManifest `json:"hlsManifests,omitempty"`
	// List of Dash manifests to use for this live encoding
	DashManifests []LiveDashManifest `json:"dashManifests,omitempty"`
	// The pass mode of the encoding
	LiveEncodingMode EncodingMode `json:"liveEncodingMode,omitempty"`
	// Reupload specific files during a live encoding. This can be helpful if an automatic life cycle policy is enabled on the output storage
	ReuploadSettings *ReuploadSettings `json:"reuploadSettings,omitempty"`
	// Major version of the manifest generator to be used for manifests referenced in this request (by properties dashManifests, dashManifests). `V2` is available for encoder versions 2.70.0 and above and is the recommended option. The default value depends on the sign-up date of your organization. See [documentation](https://developer.bitmovin.com/encoding/docs/manifest-generator-v2) page for a detailed explanation.
	ManifestGenerator ManifestGenerator `json:"manifestGenerator,omitempty"`
	// Configuration for auto restarting the live encoding
	AutoRestartConfiguration *AutoRestartConfiguration `json:"autoRestartConfiguration,omitempty"`
	// Configuration for auto shutdown of the live encoding
	AutoShutdownConfiguration *LiveAutoShutdownConfiguration `json:"autoShutdownConfiguration,omitempty"`
}
