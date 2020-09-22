package model

// EncryptionType : EncryptionType model
type EncryptionType string

// List of possible EncryptionType values
const (
	EncryptionType_AES      EncryptionType = "AES"
	EncryptionType_DE_SEDE  EncryptionType = "DESede"
	EncryptionType_BLOWFISH EncryptionType = "Blowfish"
	EncryptionType_RSA      EncryptionType = "RSA"
)
