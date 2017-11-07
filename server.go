package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func newServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	// get a default server (with logging middleware mounted)
	router := gin.Default()
	// mount routes to the server
	mountRoutes(router)
	return router
}

func mountRoutes(router *gin.Engine) {
	// mount a GET /hello/:name route to the server, which would respond with test greeting
	router.GET("/hello/:name", replyTestGreeting)
}

func replyTestGreeting(ctx *gin.Context) {
	// get parameter "name" from ctx
	name := ctx.Param("name")
	// reply with indented JSON { "Test": "Hello ${name}" }
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"Test": "Hello " + name,
	})
}
