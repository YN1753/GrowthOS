package model

import "time"

type Article struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SourceID    uint      `gorm:"index;not null" json:"source_id"`
	Title       string    `gorm:"type:varchar(512);not null" json:"title"`
	Summary     string    `gorm:"type:longtext" json:"summary"`
	Url         string    `gorm:"type:varchar(512);not null;uniqueIndex" json:"url"`
	PublishedAt time.Time `gorm:"index" json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type Source struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type"`
	Url       string    `gorm:"type:varchar(512);not null;unique" json:"url"`
	Enabled   bool      `gorm:"type:tinyint;default:1" json:"enabled"`
	Articles  []Article `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
