package model

import (
	"gorm.io/gorm"
	. "gotest/database"
	"log"
)

type Dy struct {
	gorm.Model
	SecUid       string `gorm:"primaryKey" json:"secUid"`
	UserId       string `json:"userId"`
	UserName     string `json:"userName"`
	Title        string `json:"title"`
	CollectCount int    `json:"collectCount"`
	CommentCount int    `json:"commentCount"`
	ShareCount   int    `json:"shareCount"`
	Fans         int    `json:"fans"`
	LikeCount    int32  `json:"likeCount"`
	PublicTime   string `json:"publicTime"`
}

func (dy *Dy) Create() error {
	return MyDB.Create(&dy).Error
}

func (dy *Dy) ReadAll() ([]Dy, error) {
	var dyList []Dy
	err := MyDB.Find(&dyList).Error
	return dyList, err
}

func (dy *Dy) DeleteAll() error {
	var dyList []Dy
	if err := MyDB.Find(&dyList).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := MyDB.Where("1=1").Delete(&dyList).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
