package domain

type User struct {
	Id             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DocumentTypeId int    `json:"document_type_id"`
	DocumentNumber int    `json:"document_number"`
}
