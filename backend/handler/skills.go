package handler

import (
	"net/http"
	"strconv"

	"github.com/110claw/backend/model"
	"github.com/110claw/backend/store"
	"github.com/gin-gonic/gin"
)

func ListSkills(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	tag := c.Query("tag")
	sort := c.DefaultQuery("sort", "rating")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	query := store.DB.Model(&model.Skill{}).Where("status = ?", 1)
	if tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	orderCol := "rating"
	if sort == "downloads" {
		orderCol = "download_count"
	}

	var total int64
	query.Count(&total)

	var skills []model.Skill
	query.Select("id, created_at, updated_at, name, description, version, rating, download_count, author, tags, status").
		Order(orderCol + " DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&skills)

	c.JSON(http.StatusOK, gin.H{
		"data":      skills,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func GetSkill(c *gin.Context) {
	id := c.Param("id")
	var skill model.Skill
	if err := store.DB.Where("id = ? AND status = ?", id, 1).First(&skill).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, skill)
}
