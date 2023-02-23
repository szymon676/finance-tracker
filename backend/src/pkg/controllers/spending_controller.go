package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/finance-tracker/backend/src/pkg/database"
	"github.com/szymon676/finance-tracker/backend/src/pkg/models"
)

func HandleGetSpendings(c *gin.Context) {

	var spendings []models.Spending

	database.DB.Find(&spendings)
	c.JSON(200, gin.H{
		"spendings: ": spendings,
	})
}

func HandleNewSpending(c *gin.Context) {

	var data models.BindSpending

	if err := c.BindJSON(&data); err != nil {
		c.String(400, "Error binding in spending")
	}

	//checking if user is passing good data
	if data.Title == "" || data.Price == "" || data.Description == "" {
		c.JSON(400, gin.H{
			"status": "failed, please fill all fields",
		})

	} else {

		income := models.Spending{Title: data.Title, Price: data.Price, Description: data.Description}
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

func HandleDeleteSpending(c *gin.Context) {

	var spending models.Spending
	id := c.Param("id")

	database.DB.First(&spending, id)
	if spending.Title == "" {
		c.String(404, "spending not found delete failed")
	}
	database.DB.Delete(&spending, id)
	c.String(202, "success")

}
