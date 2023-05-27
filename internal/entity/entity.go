package entity

type pet struct {
}

type user struct {
	id   int64
	name string
	pets []pet
}

type doctor struct {
	id   int64
	name string
}
