package service

import (
	"crypto/sha1"
	"fmt"
	"time"
	"errors"
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
	UserName string `json:"user_username"`
	IsFull bool `json:"user_isfull"`
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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, genereatePasswordHash(password))
	if err != nil {
		return "", err
	}
	fmt.Println("in signin user:", user)
	fmt.Println("in signin user.UserType:", user.UserType)
	fmt.Println("in signin user.isFull:", user.IsFull)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.UserType,
		user.Username,
		user.IsFull,
	})

	return token.SignedString([]byte(signingKey))
}


func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserName, nil
}


func genereatePasswordHash(password string) string{
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}