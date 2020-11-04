package model

// DolbyAtmosDialogueIntelligence : Whether to use the Dialogue Intelligence feature, which identifies and analyzes dialogue segments within audio as a basis for speech gating
type DolbyAtmosDialogueIntelligence string

// List of possible DolbyAtmosDialogueIntelligence values
const (
	DolbyAtmosDialogueIntelligence_ENABLED  DolbyAtmosDialogueIntelligence = "ENABLED"
	DolbyAtmosDialogueIntelligence_DISABLED DolbyAtmosDialogueIntelligence = "DISABLED"
)
