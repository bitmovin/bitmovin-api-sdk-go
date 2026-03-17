package model

// S3StorageClass : S3StorageClass model
type S3StorageClass string

// List of possible S3StorageClass values
const (
	S3StorageClass_GLACIER_IR          S3StorageClass = "GLACIER_IR"
	S3StorageClass_INTELLIGENT_TIERING S3StorageClass = "INTELLIGENT_TIERING"
	S3StorageClass_ONEZONE_IA          S3StorageClass = "ONEZONE_IA"
	S3StorageClass_REDUCED_REDUNDANCY  S3StorageClass = "REDUCED_REDUNDANCY"
	S3StorageClass_STANDARD            S3StorageClass = "STANDARD"
	S3StorageClass_STANDARD_IA         S3StorageClass = "STANDARD_IA"
)
