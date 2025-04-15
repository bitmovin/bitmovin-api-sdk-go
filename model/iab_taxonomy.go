package model

// IabTaxonomy model
type IabTaxonomy struct {
	Version                  *string  `json:"version,omitempty"`
	ContentTaxonomies        []string `json:"contentTaxonomies,omitempty"`
	AdOpportunityTaxonomies  []string `json:"adOpportunityTaxonomies,omitempty"`
	SensitiveTopicTaxonomies []string `json:"sensitiveTopicTaxonomies,omitempty"`
}
