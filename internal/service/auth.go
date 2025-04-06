package service

import (
	"crypto/sha1"
	"fmt"
	"todo-std"
	"todo-std/pkg/jwt"
)

const (
	salt = "qwerty123456789"
)

func (s *Service) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.userRepository.CreateUser(user)

}

func (s *Service) GenerateToken(email, password string) (string, error) {
	user, err := s.userRepository.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	return jwt.NewJWT(s.Config.Auth.Secret).Generate(user.ID)
}

func (s *Service) ParseToken(accessToken string) (int, error) {
	return jwt.NewJWT(s.Config.Auth.Secret).Parse(accessToken)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
