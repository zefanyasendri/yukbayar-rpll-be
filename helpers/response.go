package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendSuccessResponse(c *gin.Context, response Response) {
	c.JSON(http.StatusOK, response)
}

func SendCreatedResponse(c *gin.Context, response Response) {
	c.JSON(http.StatusCreated, response)
}

func SendBadRequestResponse(c *gin.Context, response Response) {
	c.JSON(http.StatusBadRequest, response)
}

func SendUnauthorzedResponse(c *gin.Context, response Response) {
	c.JSON(http.StatusBadRequest, response)
}

func SendNoContentResponse(c *gin.Context, response Response) {
	c.JSON(http.StatusNoContent, response)
}
