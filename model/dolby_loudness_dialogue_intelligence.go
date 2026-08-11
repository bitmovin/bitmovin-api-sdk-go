package model

// DolbyLoudnessDialogueIntelligence : Whether to use the Dolby Dialogue Intelligence feature, which identifies and analyzes dialogue segments within audio as a basis for speech gating
type DolbyLoudnessDialogueIntelligence string

// List of possible DolbyLoudnessDialogueIntelligence values
const (
	DolbyLoudnessDialogueIntelligence_ENABLED  DolbyLoudnessDialogueIntelligence = "ENABLED"
	DolbyLoudnessDialogueIntelligence_DISABLED DolbyLoudnessDialogueIntelligence = "DISABLED"
)
