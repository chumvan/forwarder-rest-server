package controller

import (
	"fmt"
	"net/http"

	models "github.com/chumvan/confdb/models"
	"github.com/gin-gonic/gin"
)

func MakePUTusersHandler(userChan chan []models.User) func(c *gin.Context) {
	return func(c *gin.Context) {
		var users []models.User
		if err := c.BindJSON(&users); err != nil {
			fmt.Printf("%s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "failed binding"})
			return
		}
		userChan <- users
		c.JSON(http.StatusOK, users)
	}
}
