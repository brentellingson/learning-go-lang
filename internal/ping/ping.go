// Package ping provides the implementation of the ping methods
package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents the response for the ping endpoint.
type Response struct {
	Message string `json:"message"`
}

// Controller provides the implementation of the ping endpoint.
type Controller struct{}

// NewController creates a new Ping controller.
func NewController() *Controller {
	return &Controller{}
}

// AddRoutes adds the ping routes to the router.
func (p *Controller) AddRoutes(r *gin.RouterGroup) {
	r.GET("/ping", p.Ping)
}

// Ping godoc
//
//	@Summary		ping the server
//	@Description	ping the server
//	@Tags			healthcheck
//	@Accept			"*/*"
//	@Produce		json
//	@Success		200	{object}	Response
//	@Router			/ping [get]
func (p *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "pong"})
}
