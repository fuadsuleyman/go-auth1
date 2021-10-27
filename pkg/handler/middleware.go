package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"fmt"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userName "
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		fmt.Println("Auth midleware icinde, hederParts", headerParts)
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userName, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userName)
} 