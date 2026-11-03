package model

// Visible appearance of a character. Text fields use an open vocabulary; spelling, case and unknown-value wording can vary.
type CharacterAppearance struct {
	// Summary of the character's visible appearance.
	Summary *string `json:"summary,omitempty"`
	// Apparent gender, described in free text.
	Gender *string `json:"gender,omitempty"`
	// The approximate age range of the character
	ApproximateAge AgeRange `json:"approximateAge,omitempty"`
	// Hair color.
	HairColor *string `json:"hairColor,omitempty"`
	// Hairstyle.
	HairStyle *string `json:"hairStyle,omitempty"`
	// Hair fullness or coverage.
	HairFullness *string `json:"hairFullness,omitempty"`
	// Facial hair. Can be empty when no description is available.
	FacialHair *string `json:"facialHair,omitempty"`
	// The character's physical build.
	PhysicalBuild *string `json:"physicalBuild,omitempty"`
	// Distinctive visible features, such as glasses or markings.
	DistinguishingFeatures *string `json:"distinguishingFeatures,omitempty"`
	// The character's clothing.
	Clothing *string `json:"clothing,omitempty"`
}
