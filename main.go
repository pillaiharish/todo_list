package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Email string `json:"email" binding:"required"`
}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.GET("/signup", showSignupPage)

	r.Run("localhost:10008") // listen and serve on 0.0.0.0:8080
}

func showSignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}
