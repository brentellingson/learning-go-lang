// Package main provides the entrypoint for the learning-golang-api server.
package main

import (
	"net/http"

	"github.com/brentellingson/learning-golang-api/internal/controller"
	_ "github.com/brentellingson/learning-golang-api/internal/docs"
	"github.com/brentellingson/learning-golang-api/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title									Learning Golang API
// @Version								1.0
// @Description							This is a simple API to learn Golang
// @BasePath								/api/v1
//
// @securitydefinitions.oauth2.accessCode	OAuth2Keycloak
// @authorizationUrl						http://localhost:8081/realms/myrealm/protocol/openid-connect/auth
// @tokenUrl								http://localhost:8081/realms/myrealm/protocol/openid-connect/token
func main() {
	taskService := service.NewTaskService()
	pingController := controller.NewPingController()
	taskController := controller.NewTaskController(taskService)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.Oauth2DefaultClientID("myclient")))
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})
	v1 := router.Group("/api/v1/")
	{
		pingController.AddRoutes(v1)
		taskController.AddRoutes(v1)
	}

	router.Run(":8080")
}
