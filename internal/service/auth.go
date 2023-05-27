package service

import (
	"crypto/sha1"
	"fmt"
	"pet/internal/entity"
)

const salt = "dfghnbvdcasxcdfg"

type AuthService struct {
	repo []int
}

func (auth *AuthService) CreateUser(user entity.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)

	return 0, nil
}

func (auth *AuthService) GenerateToken(user entity.User) (string, error) {
	return "", nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
