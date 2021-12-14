package model

// SimpleEncodingVodJobStatus : SimpleEncodingVodJobStatus model
type SimpleEncodingVodJobStatus string

// List of possible SimpleEncodingVodJobStatus values
const (
	SimpleEncodingVodJobStatus_CREATED   SimpleEncodingVodJobStatus = "CREATED"
	SimpleEncodingVodJobStatus_EXECUTING SimpleEncodingVodJobStatus = "EXECUTING"
	SimpleEncodingVodJobStatus_FAILURE   SimpleEncodingVodJobStatus = "FAILURE"
	SimpleEncodingVodJobStatus_RUNNING   SimpleEncodingVodJobStatus = "RUNNING"
	SimpleEncodingVodJobStatus_FINISHED  SimpleEncodingVodJobStatus = "FINISHED"
	SimpleEncodingVodJobStatus_ERROR     SimpleEncodingVodJobStatus = "ERROR"
	SimpleEncodingVodJobStatus_CANCELED  SimpleEncodingVodJobStatus = "CANCELED"
)
