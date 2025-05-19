package handlers

import (
	"more-htmx/components"

	"github.com/gin-gonic/gin"
)

var NewCount int64

func RenderCounter(c *gin.Context) {
	components.CounterExample(NewCount).Render(c.Request.Context(), c.Writer)
}

func IncrementCounter(c *gin.Context) {
	NewCount++

	components.MultiComponents(NewCount).Render(c.Request.Context(), c.Writer)
}
