package main

import (
	"fmt"

	AuthRoutes "command-line-arguments/home/user/Desktop/Golang_delivery-service_backend/Services/Auth/Application /Routes/Auth.go"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	AuthRoutes.BlogRouter(router)
	fmt.Println("Server run on port 3000")
	router.Run(":3000")
}
