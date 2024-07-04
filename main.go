package main

import (
	"fmt"
	"social-journal/initializers"
	"social-journal/routes"
	// "time"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("Social Journal App...")

	r := gin.Default()

	// CORS configuration
	// config := cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:4000"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }

	// r.Use(cors.New(config))

	routes.RegisterUserRoutes(r)
	routes.RegisterCategoryRoutes(r)
	routes.RegisterJournalRoutes(r)

	r.Run()
}
