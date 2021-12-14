package model

// SimpleEncodingVodJobAccessKeyCredentials model
type SimpleEncodingVodJobAccessKeyCredentials struct {
	// The identifier of the access key (required)
	AccessKey *string `json:"accessKey,omitempty"`
	// The secret to be used for authentication (required)
	SecretKey *string `json:"secretKey,omitempty"`
}
