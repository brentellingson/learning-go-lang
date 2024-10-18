// Package programming provides the implementation of the programming methods.
package programming

import (
	"net/http"

	"github.com/brentellingson/learning-golang-lib/programming"
	"github.com/gin-gonic/gin"
)

// Response represents the response for the UUID endpoint
type Response struct {
	UUID string `json:"uuid"`
}

// Controller provides the implementation of the programming endpoints.
type Controller struct{}

// NewController creates a new Programming controller.
func NewController() *Controller {
	return &Controller{}
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

	uuid := programming.NewUUID(withoutHyphens)
	c.JSON(http.StatusOK, Response{UUID: uuid})
}
