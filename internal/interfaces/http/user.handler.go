package http

import (
	"net/http"

	application "github.com/Alao-Abiodun/lender-api/internal/application"
	"github.com/Alao-Abiodun/lender-api/internal/domain/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (userHandler *UserHandler) Register(router *gin.Context) {
	var newUser user.User
	if err := router.ShouldBindJSON(&newUser); err != nil {
		router.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userHandler.userService.RegisterUser(router.Request.Context(), &newUser); err != nil {
		router.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	router.JSON(http.StatusCreated, gin.H{"message": " User Created successfully"})
}
