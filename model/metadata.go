package model

// Analysis metadata. Retains its original values in language-specific results.
type Metadata struct {
	// Version of the analysis software used to generate this result.
	Version *string `json:"version,omitempty"`
	// Attribution or disclaimer attached to the generated analysis.
	Disclaimer *string `json:"disclaimer,omitempty"`
}
