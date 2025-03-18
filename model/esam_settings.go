package model

// EsamSettings model
type EsamSettings struct {
	// The URL of the Placement Opportunity Information System (POIS) signal processing endpoint.  The encoder transmits SignalProcessingEvents to this endpoint whenever SCTE-35 messages are detected.
	PoisEndpoint *string `json:"poisEndpoint,omitempty"`
	// A unique identifier representing the `Acquisition Point Identity` defined in the ESAM specification.
	AcquisitionPointIdentity *string `json:"acquisitionPointIdentity,omitempty"`
	// Specifies the `Zone Identity` defined in the ESAM specification.
	ZoneIdentity *string `json:"zoneIdentity,omitempty"`
	// Defines an offset (in milliseconds) to be applied to the stream event timestamp.  This offset adjusts the `StreamTime` values (such as PTS) associated with ad opportunities  or content insertions. It is used to fine-tune timing for embedded SCTE-104/35 messages  to ensure precise frame alignment in the transport stream.
	AdAvailOffset *int32 `json:"adAvailOffset,omitempty"`
	// If authentication is required to access the POIS endpoint, credentials must be provided.
	PoisEndpointCredentials *PoisEndpointCredentials `json:"poisEndpointCredentials,omitempty"`
}
