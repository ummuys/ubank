package testutils

import "github.com/gin-gonic/gin"

func TestRouter(f gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(f)

	return r
}
