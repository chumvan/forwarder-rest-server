package router

import (
	models "github.com/chumvan/confdb/models"
	controller "github.com/chumvan/forwarder-rest-server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(usersChan chan []models.User) (r *gin.Engine) {
	r = gin.Default()

	r.PUT("/users", controller.MakePUTusersHandler(usersChan))

	return r
}
