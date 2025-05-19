package handlers

import (
	"more-htmx/components"

	"github.com/gin-gonic/gin"
)

func RenderHome(c *gin.Context) {
	components.Layout().Render(c.Request.Context(), c.Writer)
}
