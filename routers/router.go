package router

import "github.com/gin-gonic/gin"

func SetupRouter() (r *gin.Engine) {
	r = gin.Default()

	r.POST("/users")

	return
}
