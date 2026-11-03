package model

// A song identified in the video credits. Retains its original values in language-specific results.
type Song struct {
	// Song title extracted from the credit text.
	Name *string `json:"name,omitempty"`
	// Artist associated with the song in the credit text.
	Artist *string `json:"artist,omitempty"`
}
