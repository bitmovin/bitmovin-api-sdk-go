package model

// OutputType : OutputType model
type OutputType string

// List of possible OutputType values
const (
	OutputType_AKAMAI_NETSTORAGE   OutputType = "AKAMAI_NETSTORAGE"
	OutputType_AZURE               OutputType = "AZURE"
	OutputType_GENERIC_S3          OutputType = "GENERIC_S3"
	OutputType_GCS                 OutputType = "GCS"
	OutputType_FTP                 OutputType = "FTP"
	OutputType_LOCAL               OutputType = "LOCAL"
	OutputType_S3                  OutputType = "S3"
	OutputType_S3_ROLE_BASED       OutputType = "S3_ROLE_BASED"
	OutputType_SFTP                OutputType = "SFTP"
	OutputType_AKAMAI_MSL          OutputType = "AKAMAI_MSL"
	OutputType_LIVE_MEDIA_INGEST   OutputType = "LIVE_MEDIA_INGEST"
	OutputType_GCS_SERVICE_ACCOUNT OutputType = "GCS_SERVICE_ACCOUNT"
	OutputType_CDN                 OutputType = "CDN"
)
