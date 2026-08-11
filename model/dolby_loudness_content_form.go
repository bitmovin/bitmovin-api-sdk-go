package model

// DolbyLoudnessContentForm : The form of the content, used to optimize the loudness measurement gating
type DolbyLoudnessContentForm string

// List of possible DolbyLoudnessContentForm values
const (
	DolbyLoudnessContentForm_LONG        DolbyLoudnessContentForm = "LONG"
	DolbyLoudnessContentForm_SHORT       DolbyLoudnessContentForm = "SHORT"
	DolbyLoudnessContentForm_AUTO_DETECT DolbyLoudnessContentForm = "AUTO_DETECT"
)
