package model

import (
	"time"

	"gorm.io/gorm"
)

// Model is the base model embedded by all domain models.
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	Model
	Username     string  `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Phone        *string `gorm:"uniqueIndex;size:20" json:"phone"`
	PasswordHash string  `gorm:"size:255;not null" json:"-"`
	Avatar       string  `gorm:"size:500" json:"avatar"`
	Role         string  `gorm:"size:20;default:'user'" json:"role"` // user | admin
	Status       int8    `gorm:"default:1" json:"status"`            // 1=active, 0=disabled
}

type News struct {
	Model
	Title       string     `gorm:"size:200;not null" json:"title"`
	Summary     string     `gorm:"size:500" json:"summary"`
	Content     string     `gorm:"type:longtext" json:"content,omitempty"`
	CoverImage  string     `gorm:"size:500" json:"cover_image"`
	AuthorID    uint       `json:"author_id"`
	Tags        string     `gorm:"size:200" json:"tags"`
	ViewCount   int        `gorm:"default:0" json:"view_count"`
	Status      int8       `gorm:"default:1" json:"status"` // 1=published, 0=draft
	PublishedAt *time.Time `json:"published_at"`
	SourceURL   string     `gorm:"size:500;uniqueIndex" json:"source_url"` // Original URL, used for deduplication
	SourceName   string     `gorm:"size:50" json:"source_name"`              // Source: hackernews/github/reddit
}

// LearnStep represents a learning article synced from Feishu wiki.
// Previously was a 7-day plan, now stores arbitrary articles.
type LearnStep struct {
	Model
	Title       string `gorm:"size:500;not null" json:"title"`
	Description string `gorm:"size:1000" json:"description"`
	Content     string `gorm:"type:longtext" json:"content,omitempty"`
	SourceURL   string `gorm:"size:500;uniqueIndex" json:"source_url"` // Feishu wiki URL
	SortOrder   int    `gorm:"default:0" json:"sort_order"`            // Display order
	Status      int8   `gorm:"default:1" json:"status"`                // 1=active, 0=deleted
}

type Skill struct {
	Model
	Name          string  `gorm:"size:100;not null" json:"name"`
	Description   string  `gorm:"size:500" json:"description"`
	Content       string  `gorm:"type:longtext" json:"content,omitempty"`
	Version       string  `gorm:"size:50" json:"version"`
	Rating        float64 `gorm:"default:0" json:"rating"`
	DownloadCount int     `gorm:"default:0" json:"download_count"`
	Author        string  `gorm:"size:100" json:"author"`
	Tags          string  `gorm:"size:200" json:"tags"`
	Status        int8    `gorm:"default:1" json:"status"`
}

// Migrate runs AutoMigrate for all registered models.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&News{},
		&LearnStep{},
		&Skill{},
	)
}