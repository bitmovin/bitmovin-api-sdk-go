package model

// Details about user info from rtmp ingest.
type RtmpUserIngestInfo struct {
	// Client public IP address.
	Address *string `json:"address,omitempty"`
	// RTMP application name.
	App *string `json:"app,omitempty"`
	// Client stream key.
	StreamKey *string `json:"streamKey,omitempty"`
	// Flash version string / encoder identity.
	FlashVersion *string `json:"flashVersion,omitempty"`
	// Session/client connection ID.
	ClientId *string `json:"clientId,omitempty"`
	// Server action.
	EventType *string `json:"eventType,omitempty"`
}
