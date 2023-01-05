package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrBadReqHandler(err error, context *gin.Context) {
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
