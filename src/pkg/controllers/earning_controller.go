package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/szymon676/finance-tracker/src/pkg/database"
	"github.com/szymon676/finance-tracker/src/pkg/models"
)

func HandleGetEarnings(c *gin.Context) {

	var earnings []models.Earning

	database.DB.Find(&earnings)

	c.JSON(200, gin.H{
		"earnings": earnings,
	})
}

func HandleNewEarning(c *gin.Context) {

	var data models.BindEarning

	if err := c.BindJSON(&data); err != nil {
		c.String(400, "Error binding in earning")
		return
	}

	if data.Ammount == 0 || data.Category == "" || data.Currency == "" {
		c.JSON(400, gin.H{
			"status": "failed, please fill all fields",
		})
		return
	}

	layout := "02.01.2006"
	date, err := time.Parse(layout, data.Date)

	if err != nil {
		c.String(400, "error in binding date")
		return
	}

	earning := models.Earning{
		Ammount:  data.Ammount,
		Category: data.Category,
		Currency: data.Currency,
		Date:     date,
	}

	result := database.DB.Create(&earning)
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

func HandleDeleteEarning(c *gin.Context) {

	var earning models.Earning
	id := c.Param("id")

	result := database.DB.First(&earning, id)
	if result.RowsAffected == 0 {
		c.String(404, "earning not found delete failed")
		return
	}

	database.DB.Delete(&earning, id)
	c.String(202, "success")
}

func HandleGetEarningsByCategory(c *gin.Context) {
	var earnings []models.Earning
	category := c.Param("category")
	database.DB.Where("category = ?", category).Find(&earnings)
	c.JSON(200, gin.H{
		"earnings": earnings,
	})
}
