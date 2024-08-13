package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// router := gin.Default()
	// auth_routes.AuthRouter(router)
	// // orderRoutes.OrderRouter(router)

	// router.Run(":3000")

	m := gin.Default()
	m.POST("/api", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Body)
	})
	m.Run(":3000")
}
