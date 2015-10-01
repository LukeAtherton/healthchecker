package healthchecker

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() (router *gin.Engine) {
	r := gin.Default()

	gin.SetMode(gin.TestMode)

	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(200, "OK")
	})

	return r
}
