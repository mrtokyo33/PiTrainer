package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/PiTrainer/src/usecases"
)

type CreateUserHandler struct {
	useCase *usecases.CreateUserUseCase
}

func NewCreateUserHandler(useCase *usecases.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{useCase: useCase}
}

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *CreateUserHandler) CreateUser(ctx *gin.Context) {
	var body createUserRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.useCase.Execute(body.Username, body.Password); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}
