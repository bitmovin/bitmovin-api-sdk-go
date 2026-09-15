package model

// PccHdrFinding model
type PccHdrFinding struct {
	// The device pool, under the name the rest of the report shows it by. (required)
	Device     *string `json:"device,omitempty"`
	Codec      *string `json:"codec,omitempty"`
	Protection *string `json:"protection,omitempty"`
	// What was found, in one paragraph. (required)
	Sentence *string `json:"sentence,omitempty"`
}
