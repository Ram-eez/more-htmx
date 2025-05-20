package handlers

import (
	"more-htmx/components"
	"more-htmx/models"

	"github.com/gin-gonic/gin"
)

var NewCount int64
var CurrentIndex int

func RenderComponents(c *gin.Context) {
	components.Page(NewCount, 0).Render(c.Request.Context(), c.Writer)
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

var CurrentProgress int

func RenderProgressBar(c *gin.Context) {
	components.ProgressBar(CurrentProgress).Render(c.Request.Context(), c.Writer)
}

func ProgressHandler(c *gin.Context) {
	if CurrentProgress < 100 {
		CurrentProgress += 10
	}

	components.ProgressBar(CurrentProgress).Render(c.Request.Context(), c.Writer)
}

func ResetProgressBar(c *gin.Context) {
	CurrentProgress = 0
	c.Redirect(302, "/")
}
