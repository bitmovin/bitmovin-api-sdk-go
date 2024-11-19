package model

// LiveStandbyPoolEncodingStatus : Status of the standby pool encoding
type LiveStandbyPoolEncodingStatus string

// List of possible LiveStandbyPoolEncodingStatus values
const (
	LiveStandbyPoolEncodingStatus_TO_BE_CREATED LiveStandbyPoolEncodingStatus = "TO_BE_CREATED"
	LiveStandbyPoolEncodingStatus_CREATING      LiveStandbyPoolEncodingStatus = "CREATING"
	LiveStandbyPoolEncodingStatus_PREPARING     LiveStandbyPoolEncodingStatus = "PREPARING"
	LiveStandbyPoolEncodingStatus_READY         LiveStandbyPoolEncodingStatus = "READY"
	LiveStandbyPoolEncodingStatus_TO_BE_DELETED LiveStandbyPoolEncodingStatus = "TO_BE_DELETED"
	LiveStandbyPoolEncodingStatus_DELETING      LiveStandbyPoolEncodingStatus = "DELETING"
	LiveStandbyPoolEncodingStatus_ACQUIRED      LiveStandbyPoolEncodingStatus = "ACQUIRED"
	LiveStandbyPoolEncodingStatus_ERROR         LiveStandbyPoolEncodingStatus = "ERROR"
)
