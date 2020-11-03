package model


// DolbyAtmosLoudnessControl model
type DolbyAtmosLoudnessControl struct {
    // Algorithm to be used for measuring loudness. Recommended value is \"ITU_R_BS_1770_4\" (required)
    MeteringMode DolbyAtmosMeteringMode `json:"meteringMode,omitempty"`
    // Whether to use the Dialogue Intelligence feature, which identifies and analyzes dialogue segments within audio as a basis for speech gating. Must not be \"DISABLED\" for meteringMode \"ITU-R BS.1770-1\" or \"Leq (A)\", otherwise recommended value is \"ENABLED\" (required)
    DialogueIntelligence DolbyAtmosDialogueIntelligence `json:"dialogueIntelligence,omitempty"`
    // Specifies the percentage of speech that must be detected in the metered content before using the measured speech loudness as the overall program loudness. Given as an integer percentage between 0 and 100 (0% to 100%). Recommended value is 15 (required)
    SpeechThreshold *int32 `json:"speechThreshold,omitempty"`
}



