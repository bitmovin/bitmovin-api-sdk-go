package model

// Credits model
type Credits struct {
	Persons []Person `json:"persons,omitempty"`
	Songs   []Song   `json:"songs,omitempty"`
}
