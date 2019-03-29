package model

type UtcTiming struct {
	// The server to get the time from
	Value string `json:"value,omitempty"`
	// The scheme id to use. Please refer to the DASH standard.
	SchemeIdUri string `json:"schemeIdUri,omitempty"`
}

