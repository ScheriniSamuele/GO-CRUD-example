package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context, errMsg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
}

func Error404(ctx *gin.Context, errMsg string) {
	ctx.JSON(http.StatusNotFound, gin.H{"message": errMsg})
}
