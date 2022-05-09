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

func SendErrorResponse(c *gin.Context, response Response) {
	c.JSON(http.StatusBadRequest, response)
}
