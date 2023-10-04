package entity

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
