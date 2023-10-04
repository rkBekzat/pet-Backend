package entity

type Messages struct {
	Id           int
	FromUserId   int
	FromUsername string
	RoomId       int
	Content      string
	CreatedAt    string
}

type Room struct {
	Id    int `json:"id"`
	User1 int `json:"user1"`
	User2 int `json:"user2"`
}
