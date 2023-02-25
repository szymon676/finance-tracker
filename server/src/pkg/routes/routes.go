package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/finance-tracker/server/src/pkg/controllers"
	"github.com/szymon676/finance-tracker/server/src/pkg/database"
	"github.com/szymon676/finance-tracker/server/src/pkg/middlewares"
)

func SetupRoutes() {
	// * Setup routes and database
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	db := database.SetupDB
	db()

	// ! Endpoints setup for income
	r.GET("/earnings", controllers.HandleGetEarnings)
	r.POST("/earnings", controllers.HandleNewEarning)
	r.DELETE("/earnings/:id", controllers.HandleDeleteEarning)

	// * Endpoints setup for spendings
	r.GET("/spendings", controllers.HandleGetSpendings)
	r.POST("/spendings", controllers.HandleNewSpending)
	r.DELETE("/spendings/:id", controllers.HandleDeleteSpending)

	// ? category endpoints setup
	r.GET("/spendings/:category", controllers.HandleGetSpendingsByCategory)
	r.GET("/earnings/:category", controllers.HandleGetEarningsByCategory)

	// ? golas endpoints setup
	r.POST("/goals", controllers.HandleNewGoal)
	r.GET("/goals", controllers.HandleGetGoals)

	// ! running server on port 4000
	r.Run(":4000")
}
