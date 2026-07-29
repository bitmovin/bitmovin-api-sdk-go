package model

// Character model
type Character struct {
	CharacterAppearance *CharacterAppearance `json:"characterAppearance,omitempty"`
	Name                *string              `json:"name,omitempty"`
	PlayedBy            *string              `json:"playedBy,omitempty"`
	Description         *string              `json:"description,omitempty"`
}
