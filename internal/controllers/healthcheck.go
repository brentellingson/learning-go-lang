package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthcheck struct{}

type PingResponse struct {
	Message string `json:"message"`
}

func NewHealthCheckController() *healthcheck {
	return &healthcheck{}
}

func (h *healthcheck) AddRoutes(r *gin.RouterGroup) {
	r.GET("/healthcheck/ping", h.Ping)
}

// Ping godoc
//
//	@Summary		ping the server
//	@Description	ping the server
//	@Tags			healthcheck
//	@Accept			"*/*"
//	@Produce		json
//	@Success		200	{object}	PingResponse
//	@Router			/healthcheck/ping [get]
func (h *healthcheck) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, PingResponse{Message: "pong"})
}
