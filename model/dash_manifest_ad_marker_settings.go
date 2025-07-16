package model

// An optional configuration object used to define SCTE-35 signaling for ad insertion or content segmentation.  This field enables precise control over signaling behavior at both the period and event stream levels,  making it suitable for dynamic ad insertion (DAI) and live streaming workflows. This follows the ANSI/SCTE 214-1 (2024) standard.
type DashManifestAdMarkerSettings struct {
	PeriodOption          DashScte35PeriodOption          `json:"periodOption,omitempty"`
	EventStreamSignalling DashScte35EventStreamSignalling `json:"eventStreamSignalling,omitempty"`
}
