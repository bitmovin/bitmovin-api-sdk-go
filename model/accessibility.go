package model

type Accessibility struct {
	// Can be either list of languages or a complete map of services (or CC channels, in CEA-608 terminology)
	Value string `json:"value,omitempty"`
	// The scheme id to use. Please refer to the DASH standard.
	SchemeIdUri string `json:"schemeIdUri,omitempty"`
}

