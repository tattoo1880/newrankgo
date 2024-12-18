package model

import (
	"gorm.io/gorm"
	. "gotest/database"
	"log"
)

type Bil struct {
	gorm.Model
	Mid          string `gorm:"primaryKey" json:"mid"`
	UserName     string `json:"userName"`
	Title        string `json:"title"`
	Avatar       string `json:"UserHeadUrl"`
	CollectCount int    `json:"collectCount"`
	ShareCount   int    `json:"shareCount"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
	Fans         int    `json:"fans"`
	PublicTime   string `json:"publicTime"`
}

func (bil *Bil) Create() error {
	return MyDB.Create(bil).Error
}

func (bil *Bil) ReadAll() ([]Bil, error) {
	var bils []Bil
	err := MyDB.Find(&bils).Error
	return bils, err
}

func (bil *Bil) DeleteAll() error {
	var bils []Bil

	if err := MyDB.Find(&bils).Error; err != nil {
		return err
	}
	if err := MyDB.Where("1=1").Delete(&bils).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
