package model

// HlsManifestAdMarkerSettings model
type HlsManifestAdMarkerSettings struct {
	// Ad marker types that will be inserted. More than one type is possible.  - EXT_X_CUE_OUT_IN: Ad markers will be inserted using `#EXT-X-CUE-OUT` and `#EXT-X-CUE-IN` tags - EXT_OATCLS_SCTE35: Ad markers will be inserted using `#EXT-OATCLS-SCTE35` tags. They contain the base64 encoded raw bytes of the original SCTE-35 trigger. - EXT_X_SPLICEPOINT_SCTE35: Ad markers will be inserted using `#EXT-X-SPLICEPOINT-SCTE35` tags. They contain the base64 encoded raw bytes of the original SCTE-35 trigger. - EXT_X_DATERANGE: Ad markers will be inserted using `#EXT-X-DATERANGE` tags. They contain the ID, start timestamp and hex encoded raw bytes of the original SCTE-35 trigger. - EXT_X_SCTE35: Ad markers will be inserted using `#EXT-X-SCTE35` tags. They contain the base64 encoded raw bytes of the original SCTE-35 trigger.
	EnabledMarkerTypes []HlsManifestAdMarkerType `json:"enabledMarkerTypes,omitempty"`
}
