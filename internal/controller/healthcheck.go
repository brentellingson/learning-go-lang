// Package controller provides the implementation of the learning-golang-api API controllers.
package controller

import (
	"net/http"

	"github.com/brentellingson/learning-golang-api/internal/model"
	"github.com/gin-gonic/gin"
)

// HealthCheck represents the healthcheck controller.
type HealthCheck struct{}

// NewHealthCheck creates a new HealthCheck controller.
func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

// AddRoutes adds the healthcheck routes to the router.
func (h *HealthCheck) AddRoutes(r *gin.RouterGroup) {
	r.GET("/healthcheck/ping", h.Ping)
}

// Ping godoc
//
//	@Summary		ping the server
//	@Description	ping the server
//	@Tags			healthcheck
//	@Accept			"*/*"
//	@Produce		json
//	@Success		200	{object}	model.PingResponse
//	@Router			/healthcheck/ping [get]
func (h *HealthCheck) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, model.PingResponse{Message: "pong"})
}
