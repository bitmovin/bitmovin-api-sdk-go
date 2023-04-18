package model

// StreamsLiveLifeCycle : StreamsLiveLifeCycle model
type StreamsLiveLifeCycle string

// List of possible StreamsLiveLifeCycle values
const (
	StreamsLiveLifeCycle_PROVISIONING StreamsLiveLifeCycle = "PROVISIONING"
	StreamsLiveLifeCycle_STOPPED      StreamsLiveLifeCycle = "STOPPED"
	StreamsLiveLifeCycle_RUNNING      StreamsLiveLifeCycle = "RUNNING"
	StreamsLiveLifeCycle_ERROR        StreamsLiveLifeCycle = "ERROR"
)
