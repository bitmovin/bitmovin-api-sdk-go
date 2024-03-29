package model

import (
	"encoding/json"
)

// Only one of those properties may be set: fontSize, fontSizeExpression.
type TextFilter struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Can be freely chosen by the user.
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	Font       TextFilterFont          `json:"font,omitempty"`
	// If set to true a box is drawn around the text using the background color.
	Box *bool `json:"box,omitempty"`
	// The width of the box drawn around the text.
	BoxBorderWidth *int32 `json:"boxBorderWidth,omitempty"`
	// The background color to be used for drawing the box.
	BoxColor *string `json:"boxColor,omitempty"`
	// Line spacing of the border around the box in pixels
	LineSpacing *int32 `json:"lineSpacing,omitempty"`
	// Width of the border around the text
	BorderWidth *int32 `json:"borderWidth,omitempty"`
	// If set to true, it will fix text coordinates to avoid clipping if necessary
	FixBounds *bool `json:"fixBounds,omitempty"`
	// The color to be used to draw the text
	FontColor *string `json:"fontColor,omitempty"`
	// Font size to be used to draw the text
	FontSize *int32 `json:"fontSize,omitempty"`
	// An expression for the Font size. Either fontSize or fontSizeExpression can be set at the same time. The following variables are valid: main_h, h, H for input height and main_w, w, W for the input_width
	FontSizeExpression *string `json:"fontSizeExpression,omitempty"`
	// If set, alpha blending for the text is applied. Values are valid between 0.0 and 1.0.
	Alpha *float64 `json:"alpha,omitempty"`
	// Color of the shadow
	ShadowColor *string `json:"shadowColor,omitempty"`
	// X offset of the shadow
	ShadowX *int32 `json:"shadowX,omitempty"`
	// Y offset of the shadow
	ShadowY *int32 `json:"shadowY,omitempty"`
	// If set, the timecode representation in \"hh:mm:ss[:;.]ff\" format will be applied. Drop-frame timecodes (containing \";\" or \".\") must only be used with video frame rates of 29.97, 30, 59.94 or 60 FPS, according to the SMPTE standard
	Timecode *string `json:"timecode,omitempty"`
	// String to be drawn
	Text *string `json:"text,omitempty"`
	// X position of the text. Also an expression can be used. The following variables are valid: line_h - height of each text line; main_h - input height; main_w - input width; n - number of input frame; text_h - Text height; text_w - Text width (required)
	X *string `json:"x,omitempty"`
	// Y position of the text. Also an expression can be used. The following variables are valid: line_h - height of each text line; main_h - input height; main_w - input width; n - number of input frame; text_h - Text height; text_w - Text width (required)
	Y *string `json:"y,omitempty"`
	// Video frame rate
	Rate *string `json:"rate,omitempty"`
}

func (m TextFilter) FilterType() FilterType {
	return FilterType_TEXT
}
func (m TextFilter) MarshalJSON() ([]byte, error) {
	type M TextFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "TEXT"

	return json.Marshal(x)
}
