package model

// An estimated content rating. Retains its original values in language-specific results.
type Rating struct {
	// Region to which this estimated rating applies, such as US, UK or EU.
	Region *string `json:"region,omitempty"`
	// Name of the movie or television rating system used for this estimate.
	System *string `json:"system,omitempty"`
	// AI-estimated rating label within the indicated system. Can be empty when no estimate is available.
	Rating *string `json:"rating,omitempty"`
}
