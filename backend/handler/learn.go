package handler

import (
	"net/http"
	"strconv"

	"github.com/110claw/backend/model"
	"github.com/110claw/backend/store"
	"github.com/gin-gonic/gin"
)

func ListLearnSteps(c *gin.Context) {
	var steps []model.LearnStep
	store.DB.Where("status = ?", 1).
		Select("id, created_at, updated_at, day, title, description, status").
		Order("day ASC").
		Find(&steps)
	c.JSON(http.StatusOK, gin.H{"data": steps})
}

func GetLearnStep(c *gin.Context) {
	dayStr := c.Param("day")
	day, err := strconv.Atoi(dayStr)
	if err != nil || day < 1 || day > 7 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid day"})
		return
	}
	var step model.LearnStep
	if err := store.DB.Where("day = ? AND status = ?", day, 1).First(&step).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, step)
}
