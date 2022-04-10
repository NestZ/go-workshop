package models

type EmailReq struct {
	Email     string `json: "email"`
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
}
