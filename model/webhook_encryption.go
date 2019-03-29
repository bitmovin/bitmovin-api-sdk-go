package model

type WebhookEncryption struct {
	// The encryption algorithm used for the webhook
	Type EncryptionType `json:"type,omitempty"`
	// The key of the encryption
	Key string `json:"key,omitempty"`
}

