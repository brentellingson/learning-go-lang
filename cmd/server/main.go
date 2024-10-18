// Package main provides the entrypoint for the learning-golang-api server.
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/brentellingson/learning-golang-api/internal/docs"
	"github.com/brentellingson/learning-golang-api/internal/ping"
)

// @Title			Learning Golang API
// @Version		1.0
// @Description	This is a simple API to learn Golang
// @BasePath		/api/v1
func main() {
	ping := ping.NewController()

	router := gin.Default()

	docs.AddRoutes(&router.RouterGroup)
	v1 := router.Group("/api/v1/")
	{
		ping.AddRoutes(v1)
	}

	router.Run(":8080")
}
