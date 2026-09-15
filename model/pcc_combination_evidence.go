package model

// PccCombinationEvidence model
type PccCombinationEvidence struct {
	// The codec, as the shared contract spells it. (required)
	Codec *string `json:"codec,omitempty"`
	// The content protection, as the shared contract spells it. (required)
	Protection *string `json:"protection,omitempty"`
	// Device pools that played it, which is what proves the stream behind it works at all. (required)
	PlayedBy *float32 `json:"playedBy,omitempty"`
	// Device pools that reported support for it and then failed to play it. (required)
	ClaimedNotPlayedBy *float32 `json:"claimedNotPlayedBy,omitempty"`
	// Device pools that produced an answer either way. (required)
	MeasuredBy *float32 `json:"measuredBy,omitempty"`
	// No conformant stream can exist for this pairing, so it is neither gap nor result. (required)
	NotApplicable *bool `json:"notApplicable,omitempty"`
	// Hosts that served its stream. A host only — never a path and never a URL. (required)
	AssetHosts []string `json:"assetHosts,omitempty"`
	// Hosts that licensed it. A separate axis from the one above: without both, a device refusing a codec cannot be told from a stream that stopped being served. (required)
	LicenseServers []string `json:"licenseServers,omitempty"`
}
