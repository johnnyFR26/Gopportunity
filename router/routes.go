package router

import "github.com/gin-gonic/gin"

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/opening")
		v1.GET("/opening/:id")
		v1.POST("/opening")
		v1.PUT("/opening/:id")
		v1.DELETE("/opening/:id")
	}
}
