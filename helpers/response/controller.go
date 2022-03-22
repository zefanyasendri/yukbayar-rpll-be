package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(c *gin.Context, response Response) {
	c.JSON(http.StatusOK, response)
}

func SendErrorResponse(c *gin.Context, response Response) {
	c.JSON(http.StatusBadRequest, response)
}
