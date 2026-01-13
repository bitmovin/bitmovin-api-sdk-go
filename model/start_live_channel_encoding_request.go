package model

// StartLiveChannelEncodingRequest model
type StartLiveChannelEncodingRequest struct {
	// Key for the stream. (a-zA-Z, 3-20 characters) (required)
	StreamKey *string `json:"streamKey,omitempty"`
	// HLS manifests to be generated during the encoding. See [documentation](https://developer.bitmovin.com/encoding/docs/how-to-create-manifests-for-your-encodings#just-in-time-jit)
	HlsManifests []LiveHlsManifest `json:"hlsManifests,omitempty"`
	// DASH manifests to be generated during the encoding. See [documentation](https://developer.bitmovin.com/encoding/docs/how-to-create-manifests-for-your-encodings#just-in-time-jit)
	DashManifests []LiveDashManifest `json:"dashManifests,omitempty"`
	// The pass mode of the encoding. Must only be set when `encodingMode` is not set on any codec configuration used by this encoding.
	LiveEncodingMode EncodingMode `json:"liveEncodingMode,omitempty"`
	// Reupload specific files during a live encoding. This can be helpful if an automatic life cycle policy is enabled on the output storage
	ReuploadSettings *ReuploadSettings `json:"reuploadSettings,omitempty"`
	// Major version of the manifest generator to be used for manifests referenced in this request (by properties dashManifests, dashManifests). `V2` is available for encoder versions 2.70.0 and above and is the recommended option. The default value depends on the sign-up date of your organization. See [documentation](https://developer.bitmovin.com/encoding/docs/manifest-generator-v2) page for a detailed explanation.
	ManifestGenerator ManifestGenerator `json:"manifestGenerator,omitempty"`
	// Configuration for auto restarting the live encoding
	AutoRestartConfiguration *AutoRestartConfiguration `json:"autoRestartConfiguration,omitempty"`
	// Configuration for auto shutdown of the live encoding
	AutoShutdownConfiguration *LiveAutoShutdownConfiguration `json:"autoShutdownConfiguration,omitempty"`
	// Configuration for Event Signaling and Management (ESAM) system,  allowing the encoder to communicate with an ESAM server for signal processing and dynamic ad insertion update.'
	EsamSettings *EsamSettings `json:"esamSettings,omitempty"`
	// Configuration of cache control policies for media segments, HLS, and DASH manifests.  You can set caching for the HLS multivariant playlist, HLS media playlist, DASH timeline manifest,  DASH template manifest, initialization segment, and media segment.
	CacheControlSettings *CacheControlSettings `json:"cacheControlSettings,omitempty"`
	// Configuration for ad insertion features like ESAM MediaPoint insertion
	AdInsertionSettings *AdInsertionSettings `json:"adInsertionSettings,omitempty"`
}
