package models

type Person struct {
	Firstname string `json:"firstname" example:"John"`
	Lastname  string `json:"lastname" example:"Doe"`
}

type FullName struct {
	Fullname string `json:"fullname" example:"John DoeI"`
}
