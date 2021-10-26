package service

import (
	"crypto/sha1"
	"fmt"
	auth "github.com/fuadsuleyman/go-auth1"
	"github.com/fuadsuleyman/go-auth1/pkg/repository"
)

const salt = "kfjcdjkfhbsjhv6d5v1vfg"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// createuser return id or error
func (s *AuthService) CreateUser(user auth.User)(int, error){
	user.Password = genereatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func genereatePasswordHash(password string) string{
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}