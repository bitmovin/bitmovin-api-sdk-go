package model

// LiveEncodingHeartbeatUrlResponse model
type LiveEncodingHeartbeatUrlResponse struct {
	// Presigned S3 URL to download the final heartbeat JSON of the live encoding. Valid for 10 minutes. (required)
	Url *string `json:"url,omitempty"`
}
