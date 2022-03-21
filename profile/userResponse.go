package Profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/yukbayar-backend/model"
)

func sendUserSuccessresponse(c *gin.Context, ur model.PenggunaResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendUserErrorresponse(c *gin.Context, ur model.PenggunaResponse) {
	c.JSON(http.StatusBadRequest, ur)
}
