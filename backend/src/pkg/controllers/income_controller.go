package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/finance-tracker/backend/src/pkg/database"
	"github.com/szymon676/finance-tracker/backend/src/pkg/models"
)

func HandleGetIncomes(c *gin.Context) {

	var incomes []models.Income

	database.DB.Find(&incomes)
	c.JSON(200, gin.H{
		"incomes: ": incomes,
	})
}

func HandleNewIncome(c *gin.Context) {

	var data models.BindIncome

	if err := c.BindJSON(&data); err != nil {
		c.String(400, "Error binding in income")
	}

	//checking if user is passing good data
	if data.Title == "" || data.Price == "" || data.Description == "" {
		c.JSON(400, gin.H{
			"status": "failed, please fill all fields",
		})

	} else {

		income := models.Income{Title: data.Title, Price: data.Price, Description: data.Description}
		result := database.DB.Create(&income)

		//handling errors when creating user is failed
		if result.Error != nil {
			c.JSON(400, gin.H{
				"status": "failed",
			})
			return

		} else {
			// creating new user if result is successful
			c.JSON(202, gin.H{
				"status": "success",
			})

		}
	}
}

func HandleDeleteIncome(c *gin.Context) {

	var income models.Income
	id := c.Param("id")

	database.DB.First(&income, id)
	if income.Title == "" {
		c.String(404, "user not found delete failed")
	}
	database.DB.Delete(&income, id)
	c.String(202, "success")

}
