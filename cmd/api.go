package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ustadho/belajar-go-auth/handlers"
)

func main() {
	r := gin.Default()
	r.POST("/login", handlers.LoginHandler)
	r.Run("localhost:8080")
}
