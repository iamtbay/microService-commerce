package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	authService := &AuthService{}
	userRepository := &UserRepository{}
	authHandler := NewAuthHandler(userRepository, authService)
	//
	r := gin.Default()
	routesInit(r, authHandler)

	fmt.Println("Hello from users service")
	err := r.Run(":8081")
	if err != nil {
		log.Fatal("Something went wrong", err)
		os.Exit(1)
	}

}
