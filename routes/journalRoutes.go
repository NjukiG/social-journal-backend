package routes

import (
	"social-journal/controllers"
	"social-journal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterJournalRoutes(router *gin.Engine) {
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middleware.RequireAuth)

	{
		protectedRoutes.POST("/categories/:id/journals", controllers.PostJournal)
		protectedRoutes.GET("/journals", controllers.GetAllJournals)
		protectedRoutes.GET("/categories/:id/journals", controllers.GetAllJournalsByCategory)
		protectedRoutes.GET("/journals/:id", controllers.GetJournalByID)
		protectedRoutes.PUT("/journals/:id", controllers.UpdateJournal)
		protectedRoutes.DELETE("/journals/:id", controllers.DeleteJournal)

	}

}
