package handlers

import (
	"more-htmx/components"

	"github.com/gin-gonic/gin"
)

func RenderCounter(c *gin.Context) {
	components.CounterExample().Render(c.Request.Context(), c.Writer)
}
