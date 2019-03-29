package model

type XmlNamespace struct {
	// Name of the XML Namespace reference
	Prefix string `json:"prefix,omitempty"`
	// Source of the XML Namespace reference
	Uri string `json:"uri,omitempty"`
}

