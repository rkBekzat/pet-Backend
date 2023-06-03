package entity

type Pet struct {
	Id                 int     `json:"id" `
	Username           *string `json:"username"`
	Species            *string `json:"species"`
	Breed              *string `json:"breed"`
	Name               *string `json:"name"`
	DateOfBirth        *string `json:"dateOfBirth"`
	Color              *string `json:"color"`
	Sex                *bool   `json:"sex"`
	Tattoo             *string `json:"tattoo"`
	IssuedOrganization *string `json:"issued_organization"`
}

type Owner struct {
	Name        *string `json:"name"`
	AddressId   *int    `json:"address_id"`
	Surname     *string `json:"surname"`
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Username    *string `json:"username"`
	Pets        []Pet   `json:"-"`
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

type Address struct {
	AddressId     int    `json:"address_id"`
	Country       string `json:"country"`
	City          string `json:"city"`
	StateProvince string `json:"state_Province"`
	ZipPostalCode string `json:"zip_postal_code"`
	AddressLine   string `json:"address_line"`
}

type Article struct {
	ArticleId int
	Text      string
	Date      string
	AddressId int
	PhotoId   int
	UserId    int
	PetId     int
}

type Found struct {
	Article
	FoundId  int
	Expenses float32
}

type Lost struct {
	Article
	LostId int
	Reward float32
}

type ArticleComment struct {
	Text      string
	ArticleId int
}

type Photo struct {
	PhotoId int
}
