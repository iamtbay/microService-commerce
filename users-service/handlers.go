package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepository *UserRepository
	authService    *AuthService
}

func NewAuthHandler(userRepository *UserRepository, authService *AuthService) *AuthHandler {
	return &AuthHandler{authService: authService, userRepository: userRepository}
}

// GET USER
func (x *AuthHandler) GetUserHandler(c *gin.Context) {
	fmt.Println("get user")
}

// LOGIN
func (x *AuthHandler) LoginHandler(C *gin.Context) {
	fmt.Println("login")
}

// SIGN UP
func (x *AuthHandler) SignUpHandler(C *gin.Context) {
	fmt.Println("sign up")
}

// UPDATE USER
func (x *AuthHandler) UpdateUserHandler(C *gin.Context) {
	fmt.Println("update user")
}

// DELETE USER
func (x *AuthHandler) DeleteUserHandler(C *gin.Context) {
	fmt.Println("delete user")
}
