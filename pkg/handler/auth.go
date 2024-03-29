package handler

import (
	"fmt"
	"net/http"

	"github.com/fuadsuleyman/go-auth1"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	// input type
	var input auth.User

	// validasiya asagida gedir gelen data, required
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		fmt.Println("err", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func (h *Handler) signIn(c *gin.Context) {
	
	var input signInInput

	// validasiya asagida gedir gelen data, required
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		// fmt.Println("err", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":token,
	})
} 