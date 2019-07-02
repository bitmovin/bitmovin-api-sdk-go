package model

type Login struct {
	// Email address of the account. (required)
	EMail string `json:"eMail,omitempty"`
	// Password of the account. (required)
	Password string `json:"password,omitempty"`
}

