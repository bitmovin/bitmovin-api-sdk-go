package model

// A subject or object identified in a shot. Positions describe sampled appearances, not continuous tracking.
type MainSubject struct {
	// Category of the subject, usually an Open Images V7 class name such as Person or Vehicle. Other values can occur. Retains its original value in language-specific results. (required)
	Classification *string `json:"classification,omitempty"`
	// A detailed textual description of the subject's appearance and characteristics (required)
	Description *string `json:"description,omitempty"`
	// Time in seconds from the start of this shot when the subject first appears. (required)
	AppearanceTimeInSeconds *float64 `json:"appearanceTimeInSeconds,omitempty"`
	// The horizontal center position of the subject as a percentage from the left edge (0-100, where 0 is the left edge and 100 is the right edge) (required)
	CenterX *float64 `json:"centerX,omitempty"`
	// The vertical center position of the subject as a percentage from the top edge (0-100, where 0 is the top edge and 100 is the bottom edge) (required)
	CenterY *float64 `json:"centerY,omitempty"`
	// Bounding box at the subject's sampled appearance, in [ymin, xmin, ymax, xmax] order. Coordinates use a 0-1000 scale with a top-left origin, normalized independently against the full frame height and width. Omitted when no bounding box is available; the center coordinates remain valid.
	Box2d []int32 `json:"box2d,omitempty"`
}
