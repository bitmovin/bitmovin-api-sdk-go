package model

// SimpleEncodingLiveJobStatus : SimpleEncodingLiveJobStatus model
type SimpleEncodingLiveJobStatus string

// List of possible SimpleEncodingLiveJobStatus values
const (
	SimpleEncodingLiveJobStatus_CREATED   SimpleEncodingLiveJobStatus = "CREATED"
	SimpleEncodingLiveJobStatus_EXECUTING SimpleEncodingLiveJobStatus = "EXECUTING"
	SimpleEncodingLiveJobStatus_FAILURE   SimpleEncodingLiveJobStatus = "FAILURE"
	SimpleEncodingLiveJobStatus_STARTING  SimpleEncodingLiveJobStatus = "STARTING"
	SimpleEncodingLiveJobStatus_RUNNING   SimpleEncodingLiveJobStatus = "RUNNING"
	SimpleEncodingLiveJobStatus_STOPPED   SimpleEncodingLiveJobStatus = "STOPPED"
	SimpleEncodingLiveJobStatus_ERROR     SimpleEncodingLiveJobStatus = "ERROR"
	SimpleEncodingLiveJobStatus_CANCELED  SimpleEncodingLiveJobStatus = "CANCELED"
)
