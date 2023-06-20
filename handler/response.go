package handler

import (
	"fmt"
	"net/http"

	"example.com/estudoGo/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, msg string) {
	c.Header("Content-type", "aplication/json")
	c.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-type", "aplication/json")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s sucessfull ", op),
		"data":    data,
	})
}

type ErrorReponse struct {
	Message   string `json: "message"`
	ErrorCode string `json: "errorCode"`
}

type CreateOpeningReponse struct {
	Message string                  `json: "message"`
	Data    schemas.OpeningResponse `json: "data"`
}

type DeleteOpeningReponse struct {
	Message string                  `json: "message"`
	Data    schemas.OpeningResponse `json: "data"`
}

type ShowOpeningReponse struct {
	Message string                  `json: "message"`
	Data    schemas.OpeningResponse `json: "data"`
}

type ListOpeningsReponse struct {
	Message string                    `json: "message"`
	Data    []schemas.OpeningResponse `json: "data"`
}

type UpdateOpeningReponse struct {
	Message string                  `json: "message"`
	Data    schemas.OpeningResponse `json: "data"`
}
