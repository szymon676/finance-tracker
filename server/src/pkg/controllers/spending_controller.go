package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/szymon676/finance-tracker/server/src/pkg/database"
	"github.com/szymon676/finance-tracker/server/src/pkg/models"
)

func HandleGetSpendings(c *gin.Context) {
	var spendings []models.Spending
	database.DB.Find(&spendings)
	c.JSON(200, gin.H{
		"spendings": spendings,
	})
}

func HandleNewSpending(c *gin.Context) {

	var data models.BindSpending

	if err := c.BindJSON(&data); err != nil {
		c.Error(err)
		c.String(400, "Error binding in spending")
		return
	}

	//checking if user is passing good data
	if data.Ammount == 0 || data.Category == "" {
		c.JSON(400, gin.H{
			"status": "failed, please fill all fields",
		})
		return
	}

	layout := "02.01.2006"
	date, err := time.Parse(layout, data.Date)

	if err != nil {
		c.String(400, "error in date format")
		return
	}

	if date.After(time.Now()) {
		c.String(400, "date cannot be in the future")
		return
	}

	spending := models.Spending{
		Ammount:  data.Ammount,
		Category: data.Category,
		Date:     date,
	}

	result := database.DB.Create(&spending)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	c.JSON(202, gin.H{
		"status": "success",
	})
}

func HandleDeleteSpending(c *gin.Context) {

	var spending models.Spending
	id := c.Param("id")

	result := database.DB.First(&spending, id)

	if result.RowsAffected == 0 {
		c.String(404, "spending not found delete failed")
		return
	}

	database.DB.Delete(&spending, id)
	c.String(202, "success")
}

func HandleGetSpendingsByCategory(c *gin.Context) {

	var spendings []models.Spending
	category := c.Param("category")
	database.DB.Where("category = ?", category).Find(&spendings)

	c.JSON(200, gin.H{
		"spendings": spendings,
	})
}
