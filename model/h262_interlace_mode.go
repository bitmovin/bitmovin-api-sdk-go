package model
// H262InterlaceMode : Using TOP_FIELD_FIRST or BOTTOM_FIELD_FIRST will output interlaced video
type H262InterlaceMode string

// List of H262InterlaceMode
const (
	H262InterlaceMode_NONE H262InterlaceMode = "NONE"
	H262InterlaceMode_TOP_FIELD_FIRST H262InterlaceMode = "TOP_FIELD_FIRST"
	H262InterlaceMode_BOTTOM_FIELD_FIRST H262InterlaceMode = "BOTTOM_FIELD_FIRST"
)