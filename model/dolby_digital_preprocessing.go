package model

// DolbyDigitalPreprocessing model
type DolbyDigitalPreprocessing struct {
	// It indicates a gain change to be applied in the Dolby Digital decoder in order to implement dynamic range compression.  The values typically indicate gain reductions (cut) during loud passages and gain increases (boost) during quiet passages based on desired compression characteristics.
	DynamicRangeCompression *DolbyDigitalDynamicRangeCompression `json:"dynamicRangeCompression,omitempty"`
	// It applies a 120 Hz low-pass filter to the low-frequency effects (LFE) channel.  This is only allowed if the `channelLayout` contains a LFE channel.
	LfeLowPassFilter       DolbyDigitalLfeLowPassFilter       `json:"lfeLowPassFilter,omitempty"`
	NinetyDegreePhaseShift DolbyDigitalNinetyDegreePhaseShift `json:"ninetyDegreePhaseShift,omitempty"`
	ThreeDbAttenuation     DolbyDigitalThreeDbAttenuation     `json:"threeDbAttenuation,omitempty"`
}
