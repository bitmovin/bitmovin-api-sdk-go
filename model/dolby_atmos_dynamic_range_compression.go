package model

// DolbyAtmosDynamicRangeCompression model
type DolbyAtmosDynamicRangeCompression struct {
	// Line mode is intended for use in products providing line‐level or speaker‐level outputs, and is applicable to the widest range of products. Products such as set‐top boxes, DVD players, DTVs, A/V surround decoders, and outboard Dolby Atmos decoders typically use this mode.
	LineMode DolbyAtmosDynamicRangeCompressionMode `json:"lineMode,omitempty"`
	// RF mode is intended for products such as a low‐cost television receivers.
	RfMode DolbyAtmosDynamicRangeCompressionMode `json:"rfMode,omitempty"`
}
