package entity

import (
	"time"
)

type PostStatus string

const (
	Publish PostStatus = "publish"
	Draft   PostStatus = "draft"
	Trash   PostStatus = "trash"
	Unknown PostStatus = ""
)

type Post struct {
	ID        uint       `gorm:"primaryKey"`
	Title     string     `gorm:"type:varchar(200)"`
	Content   string     `gorm:"type:text"`
	Category  string     `gorm:"type:varchar(100)"`
	Status    PostStatus `gorm:"type:enum('publish', 'draft', 'trash')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
