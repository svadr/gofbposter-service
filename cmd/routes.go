package main

import (
	"github.com/gin-gonic/gin"
)

func (app *Application) routes() *gin.Engine {
	router := gin.Default()
	// Secure application against potential proxy-based attacks
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	// Handles
	router.POST("post", app.createPost)

	return router
}
