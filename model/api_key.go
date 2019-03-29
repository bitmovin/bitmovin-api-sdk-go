package model

type ApiKey struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	KeyValue string `json:"keyValue,omitempty"`
}

