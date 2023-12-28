package manager

import (
	"chat/channel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ModelAPI(c *gin.Context) {
	c.JSON(http.StatusOK, channel.ConduitInstance.GetModels())
}

func ChargeAPI(c *gin.Context) {
	c.JSON(http.StatusOK, channel.ChargeInstance.ListRules())
}

func sendErrorResponse(c *gin.Context, err error, types ...string) {
	var errType string
	if len(types) > 0 {
		errType = types[0]
	} else {
		errType = "chatnio_api_error"
	}

	c.JSON(http.StatusServiceUnavailable, TranshipmentErrorResponse{
		Error: TranshipmentError{
			Message: err.Error(),
			Type:    errType,
		},
	})
}

func abortWithErrorResponse(c *gin.Context, err error, types ...string) {
	sendErrorResponse(c, err, types...)
	c.Abort()
}
