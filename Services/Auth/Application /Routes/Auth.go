package AuthRoutes

import "github.com/gin-gonic/gin"

var BlogRouter = func(router *gin.Engine) {
	_ = router.GET("/", func(ctx *gin.Context) {
		panic("Hello world")
	})
}
