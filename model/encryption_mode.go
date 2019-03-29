package model
// EncryptionMode : The encryption method to use.
type EncryptionMode string

// List of EncryptionMode
const (
	EncryptionMode_CTR EncryptionMode = "CTR"
	EncryptionMode_CBS EncryptionMode = "CBS"
)