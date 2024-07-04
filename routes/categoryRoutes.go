package routes

import (
	"social-journal/controllers"
	"social-journal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(router *gin.Engine) {
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middleware.RequireAuth)

	{
		protectedRoutes.POST("/categories", controllers.CreateCategory)
		protectedRoutes.GET("/categories", controllers.GetAllCategories)
		protectedRoutes.GET("/categories/:id", controllers.GetCategoryByID)
		protectedRoutes.PUT("/categories/:id", controllers.UpdateCategory)
		protectedRoutes.DELETE("/categories/:id", controllers.DeleteCategory)
	}
}
