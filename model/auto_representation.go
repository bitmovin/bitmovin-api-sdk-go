package model

type AutoRepresentation struct {
	// This is the threshold if the settings of the lower or the upper representation (codec configuration) should be used, when representations are added automatically
	AdoptConfigurationThreshold *float64 `json:"adoptConfigurationThreshold,omitempty"`
}

