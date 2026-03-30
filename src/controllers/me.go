package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/PiTrainer/src/usecases"
)

type MeHandler struct {
	useCase *usecases.GetMeUseCase
}

func NewMeHandler(uc *usecases.GetMeUseCase) *MeHandler {
	return &MeHandler{useCase: uc}
}

func (h *MeHandler) GetMe(c *gin.Context) {
	userIDRaw, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found in context",
		})
		return
	}

	userID := userIDRaw.(uint)

	user, err := h.useCase.Execute(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
