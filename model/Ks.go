package model

import (
	"gorm.io/gorm"
	. "gotest/database"
	"log"
)

type Ks struct {
	gorm.Model
	UserId       string `gorm:"primaryKey" json:"userId"`
	UserName     string `json:"userName"`
	Title        string `json:"title"`
	Avatar       string `json:"UserHeadUrl"`
	LikeCount    int    `json:"likeCount"`
	ShareCount   int    `json:"shareCount"`
	CommentCount int    `json:"commentCount"`
	CollectCount int    `json:"collectCount"`
	PublicTime   string `json:"publicTime"`
	Fans         int    `json:"fans"`
}

func (ks *Ks) Create() error {
	return MyDB.Create(ks).Error
}

func (ks *Ks) ReadAll() ([]Ks, error) {
	var kss []Ks
	err := MyDB.Find(&kss).Error
	return kss, err
}

func (ks *Ks) DeleteAll() error {
	var kss []Ks

	if err := MyDB.Find(&kss).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := MyDB.Where("1=1").Delete(&kss).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
