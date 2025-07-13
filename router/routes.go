package router

import (
	"github.com/gin-gonic/gin"
	"github.com/johnnyFR26/GoOpportunity/handler"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/opening", handler.ListOpening)
		v1.GET("/opening/:id", handler.ShowOpening)
		v1.POST("/opening", handler.CreateOpening)
		v1.PUT("/opening/:id", handler.UpdateOpening)
		v1.DELETE("/opening/:id", handler.DeleteOpening)
	}
}
