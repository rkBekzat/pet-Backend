package entity

type Pet struct {
	Id                 int    `json:"id" `
	Username           string `json:"username"`
	Species            string `json:"species"`
	Breed              string `json:"breed"`
	Name               string `json:"name"`
	DateOfBirth        string `json:"dateOfBirth"`
	Color              string `json:"color"`
	Sex                bool   `json:"sex"`
	Tattoo             string `json:"tattoo"`
	IssuedOrganization string `json:"issued_organization"`
}

type Owner struct {
	Name        string `json:"name"`
	AddressId   int    `json:"address_id"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"username"`
	Pets        []Pet  `json:"-"`
}

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"-"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Doctor struct {
	Id   int64
	Name string
}
