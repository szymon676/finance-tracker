package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/szymon676/finance-tracker/src/pkg/database"
	"github.com/szymon676/finance-tracker/src/pkg/models"
)

func HandleGetGoals(c *gin.Context) {

	var goals []models.Goal
	database.DB.Find(&goals)

	c.JSON(200, gin.H{
		"goals": goals,
	})
}

func HandleNewGoal(c *gin.Context) {

	var data models.BindGoal

	if err := c.BindJSON(&data); err != nil {
		c.JSON(500, gin.H{
			"status": "error binding goal",
		})
	}

	if data.Ammount == 0 || data.Name == "" {
		c.JSON(400, gin.H{
			"message": "please pass a correct data",
		})
	}

	layout := "02.01.2006"
	date, err := time.Parse(layout, data.Date)

	if err != nil {
		c.String(400, "error in binding date")
		return
	}

	goal := models.Goal{
		Name:    data.Name,
		Ammount: data.Ammount,
		Date:    date,
	}

	result := database.DB.Create(&goal)
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

func HandleEndGoal(c *gin.Context) {

	var goal models.Goal
	id := c.Param("id")

	database.DB.Where("id = ?", id).Delete(&goal)
	c.String(202, "success")

}
