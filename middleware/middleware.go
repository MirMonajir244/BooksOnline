package middleware

import (
	"github.com/MirMonajir244/BooksOnline/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
	}
	// Extract the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header"})
		return
	}
	token := parts[1]
	tokenErr := utils.VaildateToken(token)
	if tokenErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": tokenErr.Error()})
	}

	//ctx.Set("userID",userID)
	ctx.Next()
}
