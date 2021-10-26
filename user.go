package auth

type User struct {
	Id       int    `json:"-"`
	UserType int    `json:"usertype" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
