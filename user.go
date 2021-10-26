package auth

type User struct {
	Id       int    `json:"-"`
	UserType int	`json:"usertype"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}