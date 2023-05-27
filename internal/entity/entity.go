package entity

type Pet struct {
}

type Owner struct {
	id   int64
	name string
	pets []Pet
}

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Doctor struct {
	id   int64
	name string
}
