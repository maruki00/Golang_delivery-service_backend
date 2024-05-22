package main

import (
	routes "delivery/Services/Auth/Application/Routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.AuthRouter(router)
	fmt.Println("Server run on port 3000")
	router.Run(":3000")
}
