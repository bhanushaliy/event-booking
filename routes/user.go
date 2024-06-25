package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.Users
	context.ShouldBindJSON(&user)
}
