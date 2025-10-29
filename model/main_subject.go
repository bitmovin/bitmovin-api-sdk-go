package model

// Represents a main subject or object detected within a shot, including its classification, spatial position, and appearance timing
type MainSubject struct {
	// The category or type of the detected subject based on the YOLOv8 OIv7 (Open Images V7) object detection model (e.g., 'person', 'vehicle', 'building') (required)
	Classification *string `json:"classification,omitempty"`
	// A detailed textual description of the subject's appearance and characteristics (required)
	Description *string `json:"description,omitempty"`
	// The timestamp in seconds when this subject first appears or is most prominently visible in the shot (required)
	AppearanceTimeInSeconds *float32 `json:"appearanceTimeInSeconds,omitempty"`
	// The horizontal center position of the subject as a percentage from the left edge (0-100, where 0 is the left edge and 100 is the right edge) (required)
	CenterX *float32 `json:"centerX,omitempty"`
	// The vertical center position of the subject as a percentage from the top edge (0-100, where 0 is the top edge and 100 is the bottom edge) (required)
	CenterY *float32 `json:"centerY,omitempty"`
}
