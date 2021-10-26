package auth

type User struct {
	Id       int    `json:"-"`
	UserType int	`json:"usertype"`
	Username string `json:"username"`
	Password string `json:"password"`
}