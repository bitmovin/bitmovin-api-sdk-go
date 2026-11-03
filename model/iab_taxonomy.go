package model

// IabTaxonomy model
type IabTaxonomy struct {
	// Version of the IAB Content Taxonomy used for contentTaxonomies and sensitiveTopicTaxonomies.
	ContentVersion *string `json:"contentVersion,omitempty"`
	// Version of the IAB Ad Product Taxonomy used for adOpportunityTaxonomies.
	AdProductVersion *string `json:"adProductVersion,omitempty"`
	// IAB Content Taxonomy identifiers describing the scene's subject matter. Values are identifiers, not category names. (required)
	ContentTaxonomies []string `json:"contentTaxonomies,omitempty"`
	// IAB Ad Product Taxonomy identifiers describing advertising categories relevant to this scene. Values are identifiers, not category names. (required)
	AdOpportunityTaxonomies []string `json:"adOpportunityTaxonomies,omitempty"`
	// IAB Content Taxonomy identifiers for sensitive topics associated with this scene. Values are identifiers, not category names. (required)
	SensitiveTopicTaxonomies []string `json:"sensitiveTopicTaxonomies,omitempty"`
}
