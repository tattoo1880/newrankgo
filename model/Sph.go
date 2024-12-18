package model

import (
	"gorm.io/gorm"
	. "gotest/database"
	"log"
)

type Sph struct {
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
}

func (sph *Sph) Create() error {
	return MyDB.Create(sph).Error

}

func (sph *Sph) ReadAll() ([]Sph, error) {
	var sphs []Sph
	err := MyDB.Find(&sphs).Error
	return sphs, err
}

func (sph *Sph) DeleteAll() error {
	var sphs []Sph

	if err := MyDB.Find(&sphs).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := MyDB.Where("1=1").Delete(&sphs).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
