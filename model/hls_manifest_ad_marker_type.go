package model

// HlsManifestAdMarkerType : HLS Manifest ad marker types
type HlsManifestAdMarkerType string

// List of possible HlsManifestAdMarkerType values
const (
	HlsManifestAdMarkerType_EXT_X_CUE_OUT_IN         HlsManifestAdMarkerType = "EXT_X_CUE_OUT_IN"
	HlsManifestAdMarkerType_EXT_OATCLS_SCTE35        HlsManifestAdMarkerType = "EXT_OATCLS_SCTE35"
	HlsManifestAdMarkerType_EXT_X_SPLICEPOINT_SCTE35 HlsManifestAdMarkerType = "EXT_X_SPLICEPOINT_SCTE35"
)
