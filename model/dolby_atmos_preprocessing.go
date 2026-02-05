package model

// DolbyAtmosPreprocessing model
type DolbyAtmosPreprocessing struct {
	// It indicates a gain change to be applied in the Dolby Atmos decoder in order to implement dynamic range compression.  The values typically indicate gain reductions (cut) during loud passages and gain increases (boost) during quiet passages based on desired compression characteristics.
	DynamicRangeCompression *DolbyAtmosDynamicRangeCompression `json:"dynamicRangeCompression,omitempty"`
}
