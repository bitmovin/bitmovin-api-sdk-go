package model

// IabTaxonomy model
type IabTaxonomy struct {
	ContentVersion           *string  `json:"contentVersion,omitempty"`
	AdProductVersion         *string  `json:"adProductVersion,omitempty"`
	ContentTaxonomies        []string `json:"contentTaxonomies,omitempty"`
	AdOpportunityTaxonomies  []string `json:"adOpportunityTaxonomies,omitempty"`
	SensitiveTopicTaxonomies []string `json:"sensitiveTopicTaxonomies,omitempty"`
}
