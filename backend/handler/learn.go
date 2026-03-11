package handler

import (
	"net/http"
	"strconv"

	"github.com/110claw/backend/model"
	"github.com/110claw/backend/store"
	"github.com/gin-gonic/gin"
)

func ListLearnSteps(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	var total int64
	store.DB.Model(&model.LearnStep{}).Where("status = ?", 1).Count(&total)

	var steps []model.LearnStep
	store.DB.Where("status = ?", 1).
		Select("id, created_at, updated_at, title, description, source_url, sort_order, status").
		Order("sort_order ASC, id DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&steps)

	c.JSON(http.StatusOK, gin.H{
		"data":      steps,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func GetLearnStep(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var step model.LearnStep
	if err := store.DB.Where("id = ? AND status = ?", id, 1).First(&step).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, step)
}