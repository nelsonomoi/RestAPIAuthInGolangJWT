package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/nelsonomoi/usingJWT/utils"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == ""{
			ctx.JSON(401,gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}

		err := utils.ValidateToken(tokenString)
	
		if err != nil{
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}