package model

// Character model
type Character struct {
	// Description of the character's visible appearance.
	CharacterAppearance *CharacterAppearance `json:"characterAppearance,omitempty"`
	// Name assigned to the character when identified. Can be null or an empty string, even when a summary names the person.
	Name *string `json:"name,omitempty"`
	// Name of the performer credited for the resolved character role when uniquely matched. Can be null or an empty string when identity or the role-to-performer match is unresolved. Retains its original value in language-specific results.
	PlayedBy *string `json:"playedBy,omitempty"`
	// Free-text description of the character and their role or actions in this scene.
	Description *string `json:"description,omitempty"`
}
