package routes

import (
	"more-htmx/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.Static("/static", "./static")

	router.GET("/", handlers.RenderComponents)

	router.POST("/increment", handlers.IncrementCounter)
	router.POST("/mouseclicked", handlers.ChangeClickContent)

	router.GET("/news", handlers.GetNews)
	router.GET("/progress", handlers.ProgressHandler)
	router.GET("/reset", handlers.ResetProgressBar)
	router.GET("/click", handlers.LoadButtonContent)
	router.GET("/colors", handlers.RenderColorDemo)
	router.DELETE("/fade_out_demo", handlers.DeleteDiv)

	router.POST("/restore_button", handlers.GetDeleteButton)
	router.POST("/delete_button", handlers.GetRestoreButton)

	router.GET("/refresh", handlers.OobSwap)

	router.GET("/about", handlers.AboutPage)
	router.GET("/contact", handlers.ContactPage)
}
