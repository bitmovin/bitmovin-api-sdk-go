package model


// EncryptionMode : The encryption method to use.
type EncryptionMode string

// List of possible EncryptionMode values
const (
    EncryptionMode_CTR EncryptionMode = "CTR"
    EncryptionMode_CBC EncryptionMode = "CBC"
)


