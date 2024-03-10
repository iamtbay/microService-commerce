package main

import "github.com/gin-gonic/gin"

func routesInit(r *gin.Engine, authHandler *AuthHandler)  {
	router := r.Group("/api/v1/auth")
	router.GET("/:user", authHandler.GetUserHandler)
	router.POST("/login", authHandler.LoginHandler)
	router.POST("/signup", authHandler.SignUpHandler)
	router.PUT("/edit/:userID", authHandler.UpdateUserHandler)
	router.DELETE("/:userID", authHandler.DeleteUserHandler)
}
