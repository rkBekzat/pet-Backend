package entity

type Pet struct {
	id                 int    `json:"id" `
	username           string `json:"username"`
	species            string `json:"species"`
	breed              string `json:"breed"`
	name               string `json:"name"`
	dateOfBirth        string `json:"dateOfBirth"`
	color              string `json:"color"`
	sex                bool   `json:"sex"`
	tattoo             string `json:"tattoo"`
	issuedOrganization string `json:"issued_organization"`
}

type Owner struct {
	name        string `json:"name"`
	addressId   int    `json:"address_id"`
	surname     string `json:"surname"`
	email       string `json:"email"`
	phoneNumber string `json:"phone_number"`
	username    string `json:"username"`
	pets        []Pet  `json:"-"`
}

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"-"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Doctor struct {
	id   int64
	name string
}
