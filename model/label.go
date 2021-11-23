package model

// Label model
type Label struct {
	// Identifier of the label.
	Id *int64 `json:"id,omitempty"`
	// Specifies the language of the label.
	Lang *string `json:"lang,omitempty"`
	// Content of the label. (required)
	Value *string `json:"value,omitempty"`
}
