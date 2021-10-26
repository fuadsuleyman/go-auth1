package service

import (
	auth "github.com/fuadsuleyman/go-auth1"
	"github.com/fuadsuleyman/go-auth1/pkg/repository"
)

type Authorization interface {
	CreateUser(user auth.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
