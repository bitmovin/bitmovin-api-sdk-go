package model

// CharacterAppearance model
type CharacterAppearance struct {
	Summary *string `json:"summary,omitempty"`
	Gender  *string `json:"gender,omitempty"`
	// The approximate age range of the character
	ApproximateAge         AgeRange `json:"approximateAge,omitempty"`
	HairColor              *string  `json:"hairColor,omitempty"`
	HairStyle              *string  `json:"hairStyle,omitempty"`
	HairFullness           *string  `json:"hairFullness,omitempty"`
	FacialHair             *string  `json:"facialHair,omitempty"`
	PhysicalBuild          *string  `json:"physicalBuild,omitempty"`
	DistinguishingFeatures *string  `json:"distinguishingFeatures,omitempty"`
	Clothing               *string  `json:"clothing,omitempty"`
}
