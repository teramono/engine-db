package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *DBServer) Query(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ //
		"message": "Query running...",
	})
}
