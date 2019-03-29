package model

type Login struct {
	// Email address of the account.
	EMail string `json:"eMail,omitempty"`
	// Password of the account.
	Password string `json:"password,omitempty"`
}

