package main

import (
	auth_routes "delivery/Services/Auth/Application/Routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	auth_routes.AuthRouter(router)

	router.Run(":3000")

}
