package model

// ColorConfig model
type ColorConfig struct {
	// Copy the chroma location setting from the input source
	CopyChromaLocationFlag *bool `json:"copyChromaLocationFlag,omitempty"`
	// Copy the color space setting from the input source
	CopyColorSpaceFlag *bool `json:"copyColorSpaceFlag,omitempty"`
	// Copy the color primaries setting from the input source
	CopyColorPrimariesFlag *bool `json:"copyColorPrimariesFlag,omitempty"`
	// Copy the color range setting from the input source
	CopyColorRangeFlag *bool `json:"copyColorRangeFlag,omitempty"`
	// Copy the color transfer setting from the input source
	CopyColorTransferFlag *bool `json:"copyColorTransferFlag,omitempty"`
	// The chroma location to be applied
	ChromaLocation ChromaLocation `json:"chromaLocation,omitempty"`
	// The color space to be applied. If used on a Dolby Vision stream, this value must be set to UNSPECIFIED.
	ColorSpace ColorSpace `json:"colorSpace,omitempty"`
	// The color primaries to be applied. If used on a Dolby Vision stream, this value must be set to UNSPECIFIED.
	ColorPrimaries ColorPrimaries `json:"colorPrimaries,omitempty"`
	// The color range to be applied. If used on a Dolby Vision stream, this value must be set to JPEG.
	ColorRange ColorRange `json:"colorRange,omitempty"`
	// The color transfer to be applied. If used on a Dolby Vision stream, this value must be set to UNSPECIFIED.
	ColorTransfer ColorTransfer `json:"colorTransfer,omitempty"`
	// Override the color space detected in the input file. If not set the input color space will be automatically detected if possible.
	InputColorSpace InputColorSpace `json:"inputColorSpace,omitempty"`
	// Override the color range detected in the input file. If not set the input color range will be automatically detected if possible.
	InputColorRange InputColorRange `json:"inputColorRange,omitempty"`
	// Override the color primaries detected in the input file. If not set the input color primaries will be automatically detected if possible.
	InputColorPrimaries InputColorPrimaries `json:"inputColorPrimaries,omitempty"`
	// Override the color transfer detected in the input file. If not set the input color transfer will be automatically detected if possible.
	InputColorTransfer InputColorTransfer `json:"inputColorTransfer,omitempty"`
}
