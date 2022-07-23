package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nelsonomoi/usingJWT/config"
	"github.com/nelsonomoi/usingJWT/models"
	"github.com/nelsonomoi/usingJWT/utils"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {

	var request TokenRequest 
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := config.Instance.Where("email = ?",request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := utils.GenerateJWT(user.Email,user.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}