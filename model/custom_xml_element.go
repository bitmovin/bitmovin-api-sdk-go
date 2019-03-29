package model

type CustomXmlElement struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	// String representation of the XML element
	Data string `json:"data,omitempty"`
}

