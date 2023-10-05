package repository

import (
	"github.com/jmoiron/sqlx"
	"pet/internal/entity"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessage(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

// create room and return id of room
func (m *MessagePostgres) CreateRoom(user1, user2 int) (int, error) {
	var id int
	query := "INSERT INTO Room (user1, user2) VALUES ($1, $2) RETURNING id"

	row := m.db.QueryRow(
		query,
		user1,
		user2,
	)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (m *MessagePostgres) GetRooms(id int) ([]*entity.Room, error) {
	query := "SELECT * FROM Room WHERE user1=$1 OR user2=$1"
	row, err := m.db.Query(query, id)
	if err != nil {
		return []*entity.Room{}, err
	}
	result := make([]*entity.Room, 0)
	for row.Next() {
		var res *entity.Room = new(entity.Room)
		if err = row.Scan(&res.Id, &res.User1, &res.User2); err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}

func (m *MessagePostgres) Save(message *entity.Messages) error {
	query := "INSERT INTO Messages(from_user, from_username, content, data, room_id) VALUES ($1, $2, $3, $4, $5)"

	_, err := m.db.Exec(
		query,
		message.FromUserId,
		message.FromUsername,
		message.Content,
		message.CreatedAt,
		message.RoomId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessagePostgres) GetMessages(roomId int) ([]*entity.Messages, error) {
	query := "SELECT * FROM Messages WHERE room_id=$1"
	row, err := m.db.Query(query, roomId)
	if err != nil {
		return nil, err
	}
	var result []*entity.Messages
	for row.Next() {
		var res *entity.Messages = new(entity.Messages)
		if err = row.Scan(
			&res.Id, &res.FromUserId, &res.FromUsername, &res.Content, &res.CreatedAt, &res.RoomId,
		); err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}
