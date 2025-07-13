package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
