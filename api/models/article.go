package models

import (
	"github.com/blog/db"
	"time"
)

var TArticle = Article{}

type Article struct {
	Id             int32
	Title          string
	CreatedAt      *time.Time `gorm:"column:createdAt"`
	CoverImg       string     `gorm:"column:coverImg"`
	ReadCount      int32      `gorm:"column:readCount"`
	FavouriteCount int32      `gorm:"column:favouriteCount"`
	CommentCount   int32      `gorm:"column:commentCount"`
	Abstract       string
	FilePath       string     `gorm:"column:filePath"`
	Tag            string
}

func (Article) TableName() string {
	return "article"
}

func (Article) ListArticle(pageIndex int32, pageSize int32) ([]Article, error) {
	var articles []Article
	err := db.DBHelper.DB.
		Offset(pageIndex * pageSize).
		Limit(pageSize).
		Order("createdAt desc").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
