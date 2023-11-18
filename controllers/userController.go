package controller

import "github.com/gin-gonic/gin"

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPaswword(password string) string {

}
func VerifyPassword(userPassword string, providePassword string) (bool, string) {

}
