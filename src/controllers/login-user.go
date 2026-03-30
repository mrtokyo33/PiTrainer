package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/PiTrainer/src/usecases"
)

type LoginUserHandler struct {
	useCase *usecases.LoginUserUseCase
}

func NewLoginUserHandler(uc *usecases.LoginUserUseCase) *LoginUserHandler {
	return &LoginUserHandler{useCase: uc}
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *LoginUserHandler) Login(ctx *gin.Context) {
	var body loginRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.useCase.Execute(body.Username, body.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
