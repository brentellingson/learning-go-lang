// Package programming provides the implementation of the programming methods.
package programming

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents the response for the UUID endpoint
type Response struct {
	UUID string `json:"uuid"`
}

// Service provides the programming methods interface.
type Service interface {
	NewUUID(withoutHyphens bool) string
}

// Controller provides the implementation of the programming endpoints.
type Controller struct {
	svc Service
}

// NewController creates a new Programming controller.
func NewController(svc Service) *Controller {
	return &Controller{svc}
}

// AddRoutes adds the programming routes to the router.
func (p *Controller) AddRoutes(r *gin.RouterGroup) {
	group := r.Group("/programming")
	{
		group.POST("/uuid", p.PostUUID)
	}
}

// PostUUID godoc
//
//	@Summary		create a new UUID
//	@Description	create a new UUID
//	@Tags			programming
//	@Accept			"*/*"
//	@Produce		json
//	@Param			no-hyphens	query		boolean	false	"whether to exclude hyphens in the UUID"
//	@Success		200			{object}	Response
//	@Router			/programming/uuid [post]
func (p *Controller) PostUUID(c *gin.Context) {
	withoutHyphens := c.DefaultQuery("no-hyphens", "false") == "true"

	uuid := p.svc.NewUUID(withoutHyphens)
	c.JSON(http.StatusOK, Response{UUID: uuid})
}
