package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s sucessfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json: "message"`
	Error   string `json: "error: "errorCode"`
}

type CreateOpeningResponse struct {
	Message string `json: "message"`
	Data    schemas.OpeningResponse
}

type DeleteOpeningResponse struct {
	Message string `json: "message"`
	Data    schemas.OpeningResponse
}
