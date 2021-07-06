package model

// DolbyDigitalDynamicRangeCompression model
type DolbyDigitalDynamicRangeCompression struct {
	// Line mode is intended for use in products providing line‐level or speaker‐level outputs, and is applicable to the widest range of products. Products such as set‐top boxes, DVD players, DTVs, A/V surround decoders, and outboard Dolby Digital decoders typically use this mode.
	LineMode DolbyDigitalDynamicRangeCompressionMode `json:"lineMode,omitempty"`
	// RF mode is intended for products such as a low‐cost television receivers.
	RfMode DolbyDigitalDynamicRangeCompressionMode `json:"rfMode,omitempty"`
}
