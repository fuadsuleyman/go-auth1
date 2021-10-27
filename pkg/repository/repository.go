package repository

import (
	auth "github.com/fuadsuleyman/go-auth1"
	"github.com/jmoiron/sqlx"
)

// eyni adli service.go da da interface var
type Authorization interface {
	CreateUser(user auth.User)(int, error)
	GetUser(username, password string) (auth.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}