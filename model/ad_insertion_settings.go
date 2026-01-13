package model

// AdInsertionSettings model
type AdInsertionSettings struct {
	// Enabling this allows on-demand insertion of ESAM MediaPoints. When enabled, the encoder ensures that at least one SCTE-35 data stream is available.
	EnableEsamMediaPointInsertion *bool `json:"enableEsamMediaPointInsertion,omitempty"`
	// Mode of clock synchronization during ad insertion.  If an HLS manifest is configured that has a PDT source set, the encoder defaults to the equivalent clockSynchronizationMode.  If both, HLS PDT source and clockSynchronizationMode, are set and don't match, the encoding will fail.
	ClockSynchronizationMode ClockSynchronizationMode `json:"clockSynchronizationMode,omitempty"`
}
