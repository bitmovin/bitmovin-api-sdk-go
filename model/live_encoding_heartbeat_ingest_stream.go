package model

// Details about an individual stream within the live ingest.
type LiveEncodingHeartbeatIngestStream struct {
	// Unique identifier of the stream.
	StreamId *string `json:"streamId,omitempty"`
	// Media type for the stream (e.g., \"video\" or \"audio\").
	MediaType *string `json:"mediaType,omitempty"`
	// Width of the video stream in pixels.
	Width *int32 `json:"width,omitempty"`
	// Height of the video stream in pixels.
	Height *int32 `json:"height,omitempty"`
	// Frame rate of the video stream.
	Rate *float64 `json:"rate,omitempty"`
	// Codec of the stream.
	Codec *string `json:"codec,omitempty"`
	// Aspect ratio of the video.
	AspectRatio *string `json:"aspectRatio,omitempty"`
	// Container format's bitrate of the stream, in bits per second.
	Bitrate *int32 `json:"bitrate,omitempty"`
	// Average number of samples/frames read per second.
	SamplesReadPerSecondAvg *float64 `json:"samplesReadPerSecondAvg,omitempty"`
	// Incoming bitrate measured in bits per second.
	IncomingBitrate *float64 `json:"incomingBitrate,omitempty"`
	// Largest encountered key-frame interval in milliseconds.
	KeyFrameIntervalMax *int32 `json:"keyFrameIntervalMax,omitempty"`
	// Average key-frame interval in milliseconds.
	KeyFrameIntervalAvg *float64 `json:"keyFrameIntervalAvg,omitempty"`
	// Last presentation timestamp (PTS) of the stream.
	LastTimestamp *int32 `json:"lastTimestamp,omitempty"`
	// Timescale of lastTimestamp
	LastTimestampTimescale *int32 `json:"lastTimestampTimescale,omitempty"`
	// Number of audio channels.
	NumberOfAudioChannels *int32 `json:"numberOfAudioChannels,omitempty"`
	// Format of the audio channel.
	AudioChannelFormat *string `json:"audioChannelFormat,omitempty"`
	// lastArrivalTime timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	LastArrivalTime *DateTime `json:"lastArrivalTime,omitempty"`
	// Indicates whether this particular stream is healthy.
	Healthy *bool `json:"healthy,omitempty"`
}
