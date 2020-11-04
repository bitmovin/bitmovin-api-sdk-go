package model

// Invitation model
type Invitation struct {
	Id *string `json:"id,omitempty"`
	// Email address of the tenant. (required)
	Email     *string          `json:"email,omitempty"`
	Status    InvitationStatus `json:"status,omitempty"`
	Company   *string          `json:"company,omitempty"`
	FirstName *string          `json:"firstName,omitempty"`
	LastName  *string          `json:"lastName,omitempty"`
	// Creation date of the invitation in UTC format
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	JobTitle  *string   `json:"jobTitle,omitempty"`
	Phone     *string   `json:"phone,omitempty"`
}
