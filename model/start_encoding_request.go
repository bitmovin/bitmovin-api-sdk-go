package model

// StartEncodingRequest model
type StartEncodingRequest struct {
	// Allows to encode only part of the input. Defines start (offset) and duration of the desired section. This is not allowed when the Encoding uses any kind of Input Stream resource.
	Trimming *Trimming `json:"trimming,omitempty"`
	// Scheduling parameters of the encoding.
	Scheduling *Scheduling `json:"scheduling,omitempty"`
	// Special tweaks for your encoding job.
	Tweaks *Tweaks `json:"tweaks,omitempty"`
	// Enable frame dropping/duplication to handle variable frames per seconds of video input streams
	HandleVariableInputFps *bool `json:"handleVariableInputFps,omitempty"`
	// The pass mode of the encoding. Must only be set when `encodingMode` is not set on any codec configuration used by this encoding.
	EncodingMode EncodingMode `json:"encodingMode,omitempty"`
	// DASH manifests to be generated for previewing while the encoding is still running. See [documentation](https://developer.bitmovin.com/encoding/docs/how-to-create-manifests-for-your-encodings#just-in-time-jit)
	PreviewDashManifests []ManifestResource `json:"previewDashManifests,omitempty"`
	// HLS manifests to be generated for previewing while the encoding is still running. See [documentation](https://developer.bitmovin.com/encoding/docs/how-to-create-manifests-for-your-encodings#just-in-time-jit)
	PreviewHlsManifests []ManifestResource `json:"previewHlsManifests,omitempty"`
	// DASH manifests to be generated right after encoding (just-in-time). See [documentation](https://developer.bitmovin.com/encoding/docs/how-to-create-manifests-for-your-encodings#just-in-time-jit)
	VodDashManifests []ManifestResource `json:"vodDashManifests,omitempty"`
	// HLS manifests to be generated right after encoding (just-in-time). See [documentation](https://developer.bitmovin.com/encoding/docs/how-to-create-manifests-for-your-encodings#just-in-time-jit)
	VodHlsManifests []ManifestResource `json:"vodHlsManifests,omitempty"`
	// Smooth Streaming manifests to be generated right after encoding (just-in-time). See [documentation](https://developer.bitmovin.com/encoding/docs/how-to-create-manifests-for-your-encodings#just-in-time-jit)
	VodSmoothManifests []ManifestResource `json:"vodSmoothManifests,omitempty"`
	// Major version of the manifest generator to be used for manifests referenced in this request (by properties vodDashManifests, vodHlsManifests, vodSmoothManifests, previewDashManifests, previewHlsManifests). `V2` is available for encoder versions 2.70.0 and above and is the recommended option. The default value depends on the sign-up date of your organization. See [documentation](https://developer.bitmovin.com/encoding/docs/manifest-generator-v2) page for a detailed explanation.
	ManifestGenerator ManifestGenerator `json:"manifestGenerator,omitempty"`
	// Per-Title settings
	PerTitle *PerTitle `json:"perTitle,omitempty"`
	// AI scene analysis settings
	AiSceneAnalysis *AiSceneAnalysis `json:"aiSceneAnalysis,omitempty"`
}
