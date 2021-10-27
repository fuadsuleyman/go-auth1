package repository

import (
	"fmt"

	auth "github.com/fuadsuleyman/go-auth1"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r AuthPostgres) CreateUser(user auth.User) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, usertype, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.UserType, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (auth.User, error) {
	var user auth.User
	query := fmt.Sprintf("SELECT id, usertype, username FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
} 