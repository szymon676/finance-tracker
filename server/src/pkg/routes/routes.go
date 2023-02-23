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
	r.GET("/incomes", controllers.HandleGetIncomes)
	r.POST("/incomes", controllers.HandleNewIncome)
	r.DELETE("/incomes/:id", controllers.HandleDeleteIncome)

	// * Endpoints setup for spendings
	r.GET("/spendings", controllers.HandleGetSpendings)
	r.POST("/spendings", controllers.HandleNewSpending)
	r.GET("/spendings/:id", controllers.HandleDeleteSpending)

	// ! running server on port 4000
	r.Run(":4000")
}
