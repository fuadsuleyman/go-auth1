package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	auth "github.com/fuadsuleyman/go-auth1"
	"github.com/fuadsuleyman/go-auth1/pkg/repository"
)
const (
	salt = "kfjcdjkfhbsjhv6d5v1vfg"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjHH"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	UserType int `json:"user_usertype"`
}


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

// func (s *AuthService) GenerateToken(username, password string)(string, error){
// 	user, err := s.repo.GetUser(username, genereatePasswordHash(password))
// 	if err != nil {
// 		return "", err
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodES256, &tokenClaims{
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
// 			IssuedAt: time.Now().Unix(),
// 		},
// 		user.Id,
// 		user.UserType,
// 	})
// 	return token.SignedString([]byte(signingKey))
// }

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, genereatePasswordHash(password))
	if err != nil {
		return "", err
	}
	fmt.Println("in signin user:", user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.UserType,
	})

	return token.SignedString([]byte(signingKey))
}




func genereatePasswordHash(password string) string{
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}