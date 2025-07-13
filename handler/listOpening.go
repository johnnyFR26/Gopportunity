package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
