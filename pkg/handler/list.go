package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createItem(c *gin.Context) {
	username, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"username": username,
	})
	}