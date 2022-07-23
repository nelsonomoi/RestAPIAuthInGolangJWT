package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nelsonomoi/usingJWT/models"
	"github.com/nelsonomoi/usingJWT/config"
	"log"
)

func RegisterUser(context *gin.Context)  {
	
	var user models.User 

	if err := context.ShouldBindJSON(&user);err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		context.Abort()
		log.Println("Checking here",err.Error())
		return
	}

	log.Println("request passed well")

	// fmt.Println("am here nigga")

	if err := user.HashPassword(user.Password); err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"error": err.Error})
		context.Abort()
		return
	}

	record := config.Instance.Create(&user)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})

}