package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/finance-tracker/backend/src/pkg/controllers"
	"github.com/szymon676/finance-tracker/backend/src/pkg/database"
	"github.com/szymon676/finance-tracker/backend/src/pkg/middlewares"
)

func SetupRoutes() {
	// * Setup routes and database
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	db := database.SetupDB
	db()

	// ! Endpoints setup
	r.GET("/incomes", controllers.HandleGetIncomes)
	r.POST("/incomes", controllers.HandleNewIncome)
	r.DELETE("/incomes/:id", controllers.HandleDeleteIncome)

	// ! running server on port 4000
	r.Run(":4000")
}
