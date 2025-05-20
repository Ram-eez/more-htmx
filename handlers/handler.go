package handlers

import (
	"more-htmx/components"
	"more-htmx/models"

	"github.com/gin-gonic/gin"
)

var NewCount int64
var CurrentIndex int

func RenderComponents(c *gin.Context) {
	components.Page(NewCount).Render(c.Request.Context(), c.Writer)
}

func IncrementCounter(c *gin.Context) {
	NewCount++

	components.CounterExample(NewCount).Render(c.Request.Context(), c.Writer)
}

func ChangeClickContent(c *gin.Context) {
	components.MouseClickContent().Render(c.Request.Context(), c.Writer)
}

func RenderMouseComponent(c *gin.Context) {
	components.MouseClick().Render(c.Request.Context(), c.Writer)
}

func NewsComponent(c *gin.Context) {
	components.NewsExample().Render(c.Request.Context(), c.Writer)
}

func GetNews(c *gin.Context) {
	index := CurrentIndex
	CurrentIndex = (CurrentIndex + 1) % len(models.SomeNews)

	components.NewNews(models.SomeNews, int64(index)).Render(c.Request.Context(), c.Writer)
}
