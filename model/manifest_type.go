package model


// ManifestType : ManifestType model
type ManifestType string

// List of possible ManifestType values
const (
    ManifestType_DASH ManifestType = "DASH"
    ManifestType_HLS ManifestType = "HLS"
    ManifestType_SMOOTH_STREAMING ManifestType = "SMOOTH_STREAMING"
)


