package handlers

import (
	"more-htmx/components"

	"github.com/gin-gonic/gin"
)

var NewCount int64

func RenderComponents(c *gin.Context) {
	components.CounterExample(NewCount).Render(c.Request.Context(), c.Writer)

	components.MouseClick().Render(c.Request.Context(), c.Writer)
}

func IncrementCounter(c *gin.Context) {
	NewCount++

	components.MultiComponents(NewCount).Render(c.Request.Context(), c.Writer)
}

func ChangeClickContent(c *gin.Context) {
	components.MouseClickContent().Render(c.Request.Context(), c.Writer)
}

func RenderMouseComponent(c *gin.Context) {
	components.MultiComponents(NewCount).Render(c.Request.Context(), c.Writer)
}
