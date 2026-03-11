package handler

import (
	"net/http"
	"strconv"

	"github.com/110claw/backend/model"
	"github.com/110claw/backend/store"
	"github.com/gin-gonic/gin"
)

func ListNews(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	tag := c.Query("tag")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	query := store.DB.Model(&model.News{}).Where("status = ?", 1)
	if tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	var total int64
	query.Count(&total)

	var news []model.News
	query.Select("id, created_at, updated_at, title, summary, cover_image, author_id, tags, view_count, status, published_at, source_url, source_name").
		Order("published_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&news)

	c.JSON(http.StatusOK, gin.H{
		"data":      news,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func GetNews(c *gin.Context) {
	id := c.Param("id")
	var news model.News
	if err := store.DB.Where("id = ? AND status = ?", id, 1).First(&news).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	store.DB.Model(&news).UpdateColumn("view_count", news.ViewCount+1)
	c.JSON(http.StatusOK, news)
}
