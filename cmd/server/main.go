package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/brentellingson/learning-go-lang/internal/controllers"
	_ "github.com/brentellingson/learning-go-lang/internal/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @Title			Learning Golang API
// @Version		1.0
// @Description	This is a simple API to learn Golang
// @BasePath		/api/v1
func main() {
	healthcheck := controllers.NewHealthCheckController()

	router := gin.Default()

	v1 := router.Group("/api/v1/")
	healthcheck.AddRoutes(v1)

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
