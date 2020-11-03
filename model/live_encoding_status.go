package model


// LiveEncodingStatus : LiveEncodingStatus model
type LiveEncodingStatus string

// List of possible LiveEncodingStatus values
const (
    LiveEncodingStatus_CONNECTED LiveEncodingStatus = "CONNECTED"
    LiveEncodingStatus_DISCONNECTED LiveEncodingStatus = "DISCONNECTED"
    LiveEncodingStatus_WAITING_FOR_FIRST_CONNECT LiveEncodingStatus = "WAITING_FOR_FIRST_CONNECT"
    LiveEncodingStatus_ERROR LiveEncodingStatus = "ERROR"
    LiveEncodingStatus_NOT_AVAILABLE LiveEncodingStatus = "NOT_AVAILABLE"
    LiveEncodingStatus_FINISHED LiveEncodingStatus = "FINISHED"
)


