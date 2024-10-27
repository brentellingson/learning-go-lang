package controller

import (
	"net/http"

	"github.com/brentellingson/learning-golang-api/internal/dto"
	"github.com/gin-gonic/gin"
)

// PingController provides the implementation of the ping endpoint.
type PingController struct{}

// NewPingController creates a new Ping controller.
func NewPingController() *PingController {
	return &PingController{}
}

// AddRoutes adds the ping routes to the router.
func (ctl *PingController) AddRoutes(r *gin.RouterGroup) {
	r.GET("/ping", ctl.Ping)
}

// Ping godoc
//
//	@Summary		ping the server
//	@Description	ping the server
//	@Tags			healthcheck
//	@Produce		json
//	@Success		200	{object}	dto.PingResponse
//	@Router			/ping [get]
func (ctl *PingController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, dto.PingResponse{Message: "pong"})
}
