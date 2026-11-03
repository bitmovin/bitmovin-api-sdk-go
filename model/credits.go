package model

// Credits model
type Credits struct {
	// Selected people identified in the video credits.
	Persons []Person `json:"persons,omitempty"`
	// Songs extracted from recognized credit text, such as opening, ending or insert songs. This is not a list of all music heard in the video.
	Songs []Song `json:"songs,omitempty"`
	// Production copyright year extracted from on-screen credits. Empty, null or omitted when unavailable. This is not the full copyright notice or a determination of rights or licensing. Retains its original value in language-specific results.
	Copyright *string `json:"copyright,omitempty"`
}
