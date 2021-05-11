package model

// SpriteJpegConfig model
type SpriteJpegConfig struct {
	// Quality of the JPEG file in percent. Allowed values 20 - 100 (required)
	Quality *int32 `json:"quality,omitempty"`
}
