package model

type CustomAttribute struct {
	// unique string identifier for the custom attribute
	Key string `json:"key,omitempty"`
	// value of the custom attribute
	Value string `json:"value,omitempty"`
}

