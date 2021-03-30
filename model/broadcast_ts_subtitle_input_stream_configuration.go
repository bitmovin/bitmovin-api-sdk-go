package model

// BroadcastTsSubtitleInputStreamConfiguration model
type BroadcastTsSubtitleInputStreamConfiguration struct {
	// The UUID of the subtitle stream to which this configuration belongs to. This has to be an ID of an subtitle stream that has been added to the current muxing.
	StreamId *string `json:"streamId,omitempty"`
	// An integer value. Packet Identifier (PID) for this stream.
	PacketIdentifier *int32 `json:"packetIdentifier,omitempty"`
	// The rate parameter determines the maximum rate in bits per second that should be used for the subtitle stream. The valid range is `100` to `60 000 000` bps or `0`. If the value is set to 0, we will examine the first 100 packets of subtitle packet data and use the highest rate that was computed. If the value is set too low, not enough to accommodate the subtitle bit-rate, then some PES packets corresponding to DVB subtitle stream will be dropped. This parameter is optional and the default value is 0.
	Rate *int32 `json:"rate,omitempty"`
}
